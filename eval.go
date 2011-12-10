package main

import (
	"fmt"
)

func (v cons) String() string {
	return "<cons>"
}

func doEval(c chan sexpr) {
	for e := range c {
		v := eval(global, e)
		switch v := v.(type) {
		case cons:
			fmt.Printf("<cons>\n")
		case sym:
			fmt.Printf("<sym : %s>\n", string(v))
		case float64:
			fmt.Printf("%f\n", v)
		case string:
			fmt.Printf("\"%s\"\n", v)
		case function:
			fmt.Printf("<func>\n")
		case primitive:
			fmt.Printf("<primitive>\n")
		default:
			fmt.Printf("nil\n")
		}
	}
}

func isFunction(s sexpr) bool {
	_, ok := s.(function)
	return ok
}

func isPrimitive(s sexpr) bool {
	_, ok := s.(primitive)
	return ok
}

// Perform appropriate syntax transformations on the given s-expression. Note
// that some s-expressions that 'should' involve syntax transformations, such
// as (if cond x y) and (lambda ...), don't - they just aren't evaluated as
// normal functions. (TODO make user-defined transformations more flexible to
// add symmetry.)
func transform(sc *scope, e sexpr) sexpr {
	return e
}

// Evaluates an s-expression, excluding syntax transformations (macros).
func eval(sc *scope, e sexpr) sexpr {
	e = transform(sc, e)
	switch e := e.(type) {
	case cons: // a function to evaluate
		cons := e
		car := eval(sc, cons.car)
		if !isFunction(car) && !isPrimitive(car) {
			panic("Attempted application on non-function")
		}
		cdr := cons.cdr
		args := flatten(cdr)
		if isPrimitive(car) {
			return (car.(primitive))(global, args)
		}
		f := car.(function)
		// This is a function - evaluate all arguments
		for i, a := range args {
			args[i] = eval(sc, a)
		}
		return f(global, args)
	case sym:
		return sc.lookup(string(e))
	case float64:
		return e
	case string:
		return e
	}
	return Nil
}

func flatten(s sexpr) (ss []sexpr) {
	_, ok := s.(cons)
	for ok {
		ss = append(ss, s.(cons).car)
		s = s.(cons).cdr
		_, ok = s.(cons)
	}
	// TODO what if s isn't nil now?
	return
}
