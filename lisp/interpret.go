package lisp

import (
	"fmt"
	"io"
)

func EvalFrom(ior io.Reader) {
	// TODO parse and eval in separate goroutines

	//go doParse(GetRuneScanner(r), sc)

	r := GetRuneScanner(ior)
	for {
		e := parse(r)
		v := eval(global, e)
		fmt.Println(asString(v))
	}
}
