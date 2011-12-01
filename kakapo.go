package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()

	var tc = make(chan token)
	var sc = make(chan sexpr)

	r, _ := NewPromptingReader(os.Stdin)
	go tokenize(r, tc)
	go parse(tc, sc)
	doEval(sc)
}
