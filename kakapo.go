package main

import (
	"flag"
	"os"

	. "./lisp"
)

func main() {
	flag.Parse()
	EvalFrom(os.Stdin)
}
