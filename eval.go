package main

import (
	"fmt"
	"strconv"
)

type value sexpr
type scope map[string]value
var global scope

func (v value) String() string {
	switch v.kind {
	case _ATOM:
		return v.data.(atom).String()
	case _CONS:
		return v.data.(cons).String()
	}
	panic("Reached the unreachable") // XXX
}

func (v cons) String() string {
	return "<cons>"
}

func (v atom) String() string {
	switch v.kind {
	case _NUMBER:
		return strconv.Ftoa64(v.data.(float64), 'G', -1)
	case _STRING:
		return fmt.Sprintf("\"%s\"", v.data.(string))
	case _NIL:
		return fmt.Sprintf("nil")
	case _SYMBOL:
		return v.data.(string)
	}
	return ""
}

func doEval(c chan sexpr) {
	for e := range c {
		fmt.Printf("%s\n", eval(e))
	}
}

// Evaluates an s-expression, excluding syntax transformations (macros).
func eval(e sexpr) value {
	switch e.kind {
	case _CONS:
		cons := e.data.(cons)
		car := eval(cons.car)
		cdr := cons.cdr
		// TODO
		fmt.Printf("(%s %s)\n", car, cdr)
		return car
	case _ATOM:
		a := e.data.(atom)
		switch a.kind {
		case _SYMBOL:
			return lookup(a.data.(string))
		case _NUMBER:
			return value(e)
		case _STRING:
			return value(e)
		case _NIL:
			return value(e)
		default:
			panic("Invalid kind of atom")
		}
	}
	panic("Invalid kind of sexpr")
}

// Applies the given function to an s-expression.
func apply() value {

}

// Performs lookup of symbols for values.
func lookup(sym string) value {
	// TODO attempt to lookup in reflect

	v, ok := global[sym]
	if ok {
		return v
	}
	panic("undefined")
}
