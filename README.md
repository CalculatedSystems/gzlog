# GCP + Zap Logger

Simple wrapper around [zap logger](https://github.com/uber-go/zap) with a few extra niceties for GCP logging.

## Example
Simple example with multiple outputs and formats.

```go
package main

import (
	"errors"

	log "github.com/CalculatedSystems/gzlog"
)

func init() {
	if err := log.Init(
		log.WithLevel("ERROR"),                  // set level to only ERRORS
		log.WithLogFile("/tmp/errors-only.log"), // write default output format (console) to file
		log.WithLevel("INFO"),                   // change level to INFO for subsequent options
		log.WithGCPMapping(),                    // change prefixs to match GCP logging style
		log.WithStdOut("JSON"),                  // write to stdout using JSON format
	); err != nil {
		panic(err)
	}
}

func main() {
	// writes to STDOUT only
	log.Info("giving you information")

	// support for standard library log.Print functions
	log.Println("calling Println")

	// error out to both outputs
	err := errors.New("unexpected EOF")
	log.Errorf("error during processing: %v", err)

	// writes to STDOUT and file
	log.DPanic("failure!")
}
```

**Console Output**
```
{"level":"INFO","ts":"2023-03-09T13:18:42.879-0500","msg":"giving you information"}
{"level":"INFO","ts":"2023-03-09T13:18:42.880-0500","msg":"calling Println"}
{"level":"ERROR","ts":"2023-03-09T13:18:42.880-0500","msg":"error during processing: unexpected EOF"}
{"level":"CRITICAL","ts":"2023-03-09T13:18:42.880-0500","msg":"failure!"}
```

**Logfile Contents**
```
2023-03-09T13:18:42.880-0500	error	error during processing: unexpected EOF
2023-03-09T13:18:42.880-0500	dpanic	failure!
```

Notice how the logfile shows `dpanic` while the console output displays `CRITICAL`.
