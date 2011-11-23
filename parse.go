package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type token string

func tokenize(ior io.Reader, c chan<- token) {
	// Tokenizer states
	const (
		READY = iota
		READING
		STRLIT
	)

	// Single-rune tokens
	const TOKS = "()"
	const WS = " \t\n"
	const SPLIT = TOKS + WS

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
			} else if ch == '"' {
				tmp.WriteRune(ch)
				state = STRLIT
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
		case STRLIT:
			tmp.WriteRune(ch)
			if ch == '"' {
				c <- token(tmp.String())
				tmp.Reset()
				state = READY
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
	_CONSE
)

// Types of atoms
const (
	_SYMBOLA = iota
	_NUMBERA
	_STRINGA
)

type atom struct {
	kind int
	data interface{}
}

type consE struct {
	car sexpr
	cdr sexpr
}

type sexpr struct {
	kind int
	data interface{}
}

func parse(tc <-chan token, sc chan<- sexpr) {
	// hard tokens
	const (
		LPAREN = "("
		RPAREN = ")"
	)

	for tok := range tc {
		switch tok {
		case "(":
		case ")":
			panic("Unmatched ')'")
		default:
			sc <- parseAtom(tok)
		}
	}
	close(sc)
}

func parseAtom(tok token) (e sexpr) {
	e.kind = _ATOM
	a := atom{}

	a.kind = _SYMBOLA
	a.data = string(tok)

	// try as string literal
	if tok[0] == '"' {
		a.kind = _STRINGA
		a.data = string(tok[1:len(tok)-1])
	}

	// try as number
	n, err := strconv.Atof64(string(tok))
	if err == nil {
		a.kind = _NUMBERA
		a.data = n
	}

	e.data = a
	return
}
