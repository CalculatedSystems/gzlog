package log

import (
	"errors"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// l is the global zap.SugaredLogger. Any calls to pacakge level functions
	// use this logger under the hood.
	l *zap.SugaredLogger

	cfg *setupConfig
)

type setupOpt func(*setupConfig) error

type setupConfig struct {
	cores []zapcore.Core
	zap   zap.Config
}

func init() {
	// establish a basic logger on init to prevent panic(s)
	p, _ := zap.NewProduction()
	l = p.Sugar()

	cfg = &setupConfig{
		cores: make([]zapcore.Core, 0),
		zap:   zap.NewProductionConfig(),
	}
	cfg.zap.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
}

func Logger() *zap.SugaredLogger {
	return l
}

// Init modifies the global logger based on options provided.
// Errors on any issue provided by options
func Init(opts ...setupOpt) (err error) {
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return err
		}
	}

	l = zap.New(zapcore.NewTee(cfg.cores...)).Sugar()

	return nil
}

func WithLogFile(filename string) setupOpt {
	return func(cfg *setupConfig) error {
		// create|open log file
		logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(cfg.zap.EncoderConfig),
			zapcore.AddSync(logfile),
			cfg.zap.Level,
		)

		cfg.cores = append(cfg.cores, core)

		return nil
	}
}

func WithLevel(level string) setupOpt {
	return func(cfg *setupConfig) error {
		al, err := zap.ParseAtomicLevel(level)
		if err != nil {
			return err
		}
		cfg.zap.Level = al

		return nil
	}
}

func WithStdOut(encoding string) setupOpt {
	return func(cfg *setupConfig) error {
		var encoder zapcore.Encoder

		switch strings.ToUpper(encoding) {
		case "JSON":
			encoder = zapcore.NewJSONEncoder(cfg.zap.EncoderConfig)
		case "CONSOLE":
			encoder = zapcore.NewConsoleEncoder(cfg.zap.EncoderConfig)
		default:
			return errors.New("unknown encoding")
		}

		cfg.cores = append(cfg.cores, zapcore.NewCore(encoder, os.Stdout, cfg.zap.Level))

		return nil
	}
}

func WithGCPMapping() setupOpt {
	return func(cfg *setupConfig) error {
		cfg.zap.EncoderConfig.TimeKey = "time"
		cfg.zap.EncoderConfig.LevelKey = "severity"
		cfg.zap.EncoderConfig.NameKey = "logger"
		cfg.zap.EncoderConfig.CallerKey = "caller"
		cfg.zap.EncoderConfig.MessageKey = "message"
		cfg.zap.EncoderConfig.StacktraceKey = "stacktrace"
		cfg.zap.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
		cfg.zap.EncoderConfig.EncodeLevel = gcpEncodeLevel()
		cfg.zap.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		cfg.zap.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
		cfg.zap.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		return nil
	}
}

// gcpEncodeLevel provided by https://github.com/uber-go/zap/discussions/1110
func gcpEncodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}
