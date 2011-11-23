package main

import (
	"fmt"
	"strconv"
)

// Types of values
const (
	_STRING = iota
	_NUMBER
	_CONS
)

type value struct {
	kind int
	data interface{}
}

func (v value) String() string {
	switch v.kind {
	case _NUMBER:
		return strconv.Ftoa64(v.data.(float64), 'G', -1)
	case _STRING:
		return fmt.Sprintf("\"%s\"", v.data.(string))
	}
	return ""
}

func doEval(c chan sexpr) {
	for e := range c {
		fmt.Printf("%s\n", eval(e))
	}
}

func eval(e sexpr) value {
	switch e.kind {
	case _CONSE:
		fmt.Println("Cons")
	case _ATOM:
		v := value{}
		a := e.data.(atom)
		switch a.kind {
		case _SYMBOLA:
		case _NUMBERA:
			v.kind = _NUMBER
			v.data = a.data
		case _STRINGA:
			v.kind = _STRING
			v.data = a.data
		default:
			panic("Invalid kind of atom")
		}
		return v
	default:
		panic("Invalid kind of sexpr")
	}
	panic("Reached the unreachable") // XXX
}
