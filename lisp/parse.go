package main

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

type token string

func tokenize(ior io.Reader, c chan<- token) {
	// Tokenizer states
	const (
		READY = iota
		READING
		STRLIT
		ESCAPE
		COMMENT
	)

	// Single-rune tokens
	const TOKS = "()"
	const WS = " \t\n"
	const SPLIT = TOKS + WS + ";"

	r := bufio.NewReader(ior)

	state := READY
	var tmp bytes.Buffer

	ch, _, err := r.ReadRune()
	for err == nil {
		switch state {
		case READY:
			// c either begins or is a token
			if strings.ContainsRune(TOKS, ch) {
				c <- token(ch)
			} else if strings.ContainsRune(WS, ch) {
				// whitespace; ignore it
			} else if ch == ';' {
				// read to EOL
				state = COMMENT
			} else if ch == '"' {
				tmp.WriteRune(ch)
				state = STRLIT
			} else {
				tmp.WriteRune(ch)
				state = READING
			}
		case READING:
			if strings.ContainsRune(SPLIT, ch) {
				// the current token is done
				c <- token(tmp.String())
				tmp.Reset()
				state = READY
				r.UnreadRune()
			} else {
				tmp.WriteRune(ch)
			}
		case STRLIT:
			if ch == '\\' {
				state = ESCAPE
			} else {
				tmp.WriteRune(ch)
				if ch == '"' {
					c <- token(tmp.String())
					tmp.Reset()
					state = READY
				}
			}
		case ESCAPE:
			switch ch {
			case 'n':
				tmp.WriteRune('\n')
			case 't':
				tmp.WriteRune('\t')
			default:
				panic("Invalid escape character")
			}
			state = STRLIT
		case COMMENT:
			if ch == '\n' {
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

type sexpr interface{}
type atom interface{}
type sym string

type cons struct {
	car sexpr
	cdr sexpr
}

var Nil = interface{}(nil)

// hard tokens
const (
	_LPAREN = "("
	_RPAREN = ")"
)

func parse(tc <-chan token, sc chan<- sexpr) {
	for tok := range tc {
		sc <- parseNext(tok, tc)
	}
	close(sc)
}

func parseNext(tok token, tc <-chan token) sexpr {
	switch tok {
	case _LPAREN:
		return parseCons(tc)
	case _RPAREN:
		panic("Unmatched ')'")
	}
	return parseAtom(tok)
}

func parseCons(tc <-chan token) sexpr {
	// note that we assume the LPAREN has already been read
	tok := <-tc
	if tok == _RPAREN {
		// nil atom
		return Nil
	}
	car := parseNext(tok, tc)
	cdr := parseCons(tc)
	return cons{car, cdr}
}

func parseAtom(tok token) (e sexpr) {
	e = sym(tok)

	// try as string literal
	if tok[0] == '"' {
		e = string(tok[1 : len(tok)-1])
	}

	// try as number
	n, err := strconv.ParseFloat(string(tok), 64)
	if err == nil {
		e = n
	}
	return
}
