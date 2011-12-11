package lisp

import (
	"io"
)

func EvalFrom(ior io.Reader) {
	// TODO parse and eval in separate goroutines

	//go doParse(GetRuneScanner(r), sc)

	r := GetRuneScanner(ior)
	e, err := parse(r)
	for err == nil {
		eval(global, e)
		e, err = parse(r)
	}
}
