package util

import (
	"log"
	"time"
)

// Elapsed measures the execution time.
// it can be used via defer util.Elapsed("functionname")
// See also: https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", what, time.Since(start))
	}
}
