package lisp

import (
	"bufio"
	"io"
	"strings"
)

func load(s string) {
	EvalFrom(strings.NewReader(s))
}

func EvalFrom(ior io.Reader) {
	// TODO parse and eval in separate goroutines

	r := bufio.NewReader(ior)
	e, err := parse(r)
	for err == nil {
		eval(global, e)
		e, err = parse(r)
	}
}
