package main

import (
	"flag"
	"strings"

	. "./lisp"
)

func main() {
	flag.Parse()
	EvalFrom(strings.NewReader(repl))
}
