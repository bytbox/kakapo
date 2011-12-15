package lisp

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type token string

func readToken(r io.RuneScanner) (token, error) {
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
	const WS = " \t\r\n"
	const SPLIT = TOKS + WS + ";"
	const PROTECT = '\''

	state := READY
	var tmp bytes.Buffer

	ch, _, err := r.ReadRune()
	if err != nil {
		return "", err
	}
	for err == nil {
		switch state {
		case READY:
			// c either begins or is a token
			if strings.ContainsRune(TOKS, ch) {
				return token(ch), nil
			} else if strings.ContainsRune(WS, ch) {
				// whitespace; ignore it
			} else if ch == ';' {
				// read to EOL
				state = COMMENT
			} else if ch == '"' {
				tmp.WriteRune(ch)
				state = STRLIT
			} else if ch == PROTECT {
				return token(ch), nil
			} else {
				tmp.WriteRune(ch)
				state = READING
			}
		case READING:
			if strings.ContainsRune(SPLIT, ch) {
				// the current token is done
				tok := token(tmp.String())
				tmp.Reset()
				state = READY
				r.UnreadRune()
				return tok, nil
			} else {
				tmp.WriteRune(ch)
			}
		case STRLIT:
			if ch == '\\' {
				state = ESCAPE
			} else {
				tmp.WriteRune(ch)
				if ch == '"' {
					tok := token(tmp.String())
					tmp.Reset()
					state = READY
					return tok, nil
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

	if err != io.EOF {
		return "", err
	}
	switch state {
	case READY:
		return "", err
	case COMMENT:
		return "", err
	}
	panic("Unexpected EOF")
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
	_LPAREN  = "("
	_RPAREN  = ")"
	_PROTECT = "'"
)

func parse(r io.RuneScanner) (sexpr, error) {
	tok, err := readToken(r)
	if err == nil {
		return parseNext(tok, r), nil
	}
	return Nil, err
}

func parseNext(tok token, r io.RuneScanner) sexpr {
	switch tok {
	case _LPAREN:
		return parseCons(r)
	case _RPAREN:
		panic("Unmatched ')'")
	case _PROTECT:
		s, e := parse(r)
		if e != nil {
			panic(e)
		}
		return cons{sym("quote"), cons{s, nil}}
	}
	return parseAtom(tok)
}

func parseCons(r io.RuneScanner) sexpr {
	// note that we assume the LPAREN has already been read
	tok, err := readToken(r)
	if err != nil {
		panic(err)
	}
	if tok == _RPAREN {
		// nil atom
		return Nil
	}
	if tok == "." {
		tok, err := readToken(r)
		if err != nil {
			panic(err)
		}
		ret := parseNext(tok, r)
		tok, err = readToken(r)
		if err != nil {
			panic(err)
		}
		if tok != _RPAREN {
			panic("Expected ')'")
		}
		return ret
	}
	car := parseNext(tok, r)
	cdr := parseCons(r)
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
