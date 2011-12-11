package main

import (
	"io"
)

func EvalFrom(r io.Reader) {
	var tc = make(chan token)
	var sc = make(chan sexpr)

	go tokenize(r, tc)
	go parse(tc, sc)
	doEval(sc)
}
