package main

import (
	"flag"
	"fmt"
	"os"
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

	ExposeGlobal("-interpreter", "Kakapo")
	ExposeGlobal("-interpreter-version", VERSION)

	if len(flag.Args()) > 0 {
		for _, fname := range flag.Args() {
			f, err := os.Open(fname)
			if err != nil {
				panic(err)
			}
			EvalFrom(f)
		}
		return
	}

	// Start the read-eval-print loop
	EvalFrom(strings.NewReader(repl))
}
