package main

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/CalculatedSystems/gzlog"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	level := "DEBUG"
	message := "standard message"

	params := r.URL.Query()
	if params.Has("level") {
		level = params.Get("level")
	}
	if params.Has("message") {
		message = params.Get("message")
	}

	l := strings.ToUpper(level)
	switch {
	case l == "ERROR":
		err := errors.New("im a golang error type")
		log.Errorw(message, "err", err)
	case l == "WARN":
		log.Warn(message)
	case l == "DEBUG":
		log.Debug(message)
	case l == "INFO":
		log.Info(message)
	case strings.Contains(l, "PANIC") || strings.Contains(l, "FATAL"):
		log.Warnw("user tried to cause panic", "req", r)
		http.Error(w, "nice try, but no dice", http.StatusBadRequest)
	default:
		http.Error(w, "unknown level", http.StatusBadRequest)
	}
}

func init() {
	if err := log.Init(
		log.WithLevel("DEBUG"), // change level to INFO for subsequent options
		log.WithGCPMapping(),   // change prefixs to match GCP logging style
		log.WithStdOut("JSON"), // write to stdout using JSON format
	); err != nil {
		panic(err)
	}
}

func main() {
	log.Info("starting up...")
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
