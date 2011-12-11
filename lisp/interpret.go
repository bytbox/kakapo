package lisp

import (
	"io"
)

func EvalFrom(r io.Reader) {
	var sc = make(chan sexpr)

	go doParse(GetRuneScanner(r), sc)
	doEval(sc)
}
