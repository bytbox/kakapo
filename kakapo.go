package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	. "github.com/bytbox/kakapo/lisp"
)

const VERSION = `0.4`

var (
	version = flag.Bool("V", false, "Display version information and exit")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Kakapo %s\n", VERSION)
		return
	}

	// Expose imports
	for name, pkg := range _go_imports {
		ExposeImport(name, pkg)
	}

	// Expose globals
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

	args := flag.Args()
	if len(args) == 0 {
		// Start the read-eval-print loop
		EvalFrom(strings.NewReader(repl))
	} else {
		for _, path := range args {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			EvalFrom(file)
		}
	}
}

