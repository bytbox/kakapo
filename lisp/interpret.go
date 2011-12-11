package lisp

import (
	"bufio"
	"io"
)

func EvalFrom(r io.Reader) {
	var sc = make(chan sexpr)

	go parse(bufio.NewReader(r), sc)
	doEval(sc)
}
