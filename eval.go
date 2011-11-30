package main

import (
	"fmt"
	"strconv"
)

type scope map[string]sexpr

func (v sexpr) String() string {
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

func isFunction(s sexpr) bool {
	if s.kind != _ATOM {
		return false
	}
	a := s.data.(atom)
	if a.kind != _FUNCTION {
		return false
	}
	return true
}

func isSyntax(s sexpr) bool {
	return false
}

// Perform appropriate syntax transformations on the given s-expression. Note
// that some s-expressions that 'should' involve syntax transformations, such
// as (if cond x y) and (lambda ...), don't - they just aren't evaluated as
// normal functions. (TODO make user-defined transformations more flexible to
// add symmetry.)
func transform(e sexpr) sexpr {
	return e // TODO
}

// Evaluates an s-expression, excluding syntax transformations (macros).
func eval(e sexpr) sexpr {
	switch e.kind {
	case _CONS: // a function to evaluate
		cons := e.data.(cons)
		car := eval(cons.car)
		cdr := cons.cdr
		if !isFunction(car) && !isSyntax(car) {
			panic("Attempted application on non-function")
		}
		if isFunction(car) {
			return apply(car.data.(atom).data.(func([]sexpr) sexpr), cdr)
		} else { // isSyntax(car)
			return Nil // TODO
		}
	case _ATOM:
		a := e.data.(atom)
		switch a.kind {
		case _SYMBOL:
			return lookup(a.data.(string))
		case _NUMBER:
			return e
		case _STRING:
			return e
		case _NIL:
			return e
		default:
			panic("Invalid kind of atom")
		}
	}
	panic("Invalid kind of sexpr")
}

func flatten(s sexpr) (ss []sexpr) {
	for s.kind == _CONS {
		ss = append(ss, s.data.(cons).car)
		s = s.data.(cons).cdr
	}
	// TODO what if s isn't nil now?
	return
}

// Applies the given function to an s-expression.
func apply(f func ([]sexpr) sexpr, args sexpr) sexpr {
	return f(flatten(args))
}

// Performs lookup of symbols for values.
func lookup(sym string) sexpr {
	// TODO attempt lookup in reflect

	v, ok := global[sym]
	if ok {
		return v
	}
	panic("undefined")
}
