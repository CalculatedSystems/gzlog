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
