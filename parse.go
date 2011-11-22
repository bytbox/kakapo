package main

import (
	"bufio"
	"bytes"
	"io"
)

type token string

func tokenize(ior io.Reader, c chan<- token) {
	// Tokenizer states
	const (
		READY = iota
		READING
	)

	// Single-rune tokens
	const TOKS = "()"
	const WS = " \t\n"
	const SPLIT = TOKS + WS + `"`

	r := bufio.NewReader(ior)

	state := READY
	var tmp bytes.Buffer

	ch, _, err := r.ReadRune()
	for err == nil {
		switch state {
		case READY:
			// c either begins or is a token
			if ContainsRune(TOKS, ch) {
				c <- token(ch)
			} else if ContainsRune(WS, ch) {
				// whitespace; ignore it
			} else {
				tmp.WriteRune(ch)
				state = READING
			}
		case READING:
			if ContainsRune(SPLIT, ch) {
				// the current token is done
				c <- token(tmp.String())
				tmp.Reset()
				state = READY
				r.UnreadRune()
			} else {
				tmp.WriteRune(ch)
			}
		default:
			panic("Invalid state")
		}
		ch, _, err = r.ReadRune()
	}

	close(c)
	if err != nil && err != io.EOF {
		panic(err)
	}
}

// Types of s-expressions
const (
	_ATOM = iota
	_CONS
)

// Types of atoms
const (
	_SYMBOL = iota
	_NUMBER
	_STRING
)

type atom struct {
	kind int
	data interface{}
}

type cons struct {
	car sexpr
	cdr sexpr
}

type sexpr struct {
	kind int
	data interface{}
}

func parse(tc <-chan token, sc chan<- sexpr) {
	for tok := range tc {
		println(tok)
	}
	close(sc)
}
