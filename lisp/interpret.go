package lisp

import (
	"fmt"
	"io"
)

func EvalFrom(ior io.Reader) {
	// TODO parse and eval in separate goroutines

	//go doParse(GetRuneScanner(r), sc)

	r := GetRuneScanner(ior)
	e, err := parse(r)
	for err == nil {
		v := eval(global, e)
		fmt.Println(asString(v))
		e, err = parse(r)
	}
}
