package main

import (
	"flag"
	"fmt"
	"strings"

	. "./lisp"
)

const VERSION = `0.2`

var (
	version = flag.Bool("V", false, "Display version information and exit")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Kakapo %s\n", VERSION)
		return
	}

	// Start the read-eval-print loop
	EvalFrom(strings.NewReader(repl))
}
