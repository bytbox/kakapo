package lisp

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
			return (car.(primitive))(sc, args)
		}
		f := car.(function)
		// This is a function - evaluate all arguments
		for i, a := range args {
			args[i] = eval(sc, a)
		}
		return f(sc, args)
	case sym:
		return sc.lookup(e)
	}
	return e
}

func apply(sc *scope, e sexpr, ss []sexpr) sexpr {
	f, ok := e.(function)
	if !ok {
		panic("Attempted application on non-function")
	}
	return f(sc, ss)
}

func unflatten(ss []sexpr) sexpr {
	c := sexpr(nil)
	for i := len(ss) - 1; i >= 0; i-- {
		c = cons{ss[i], c}
	}
	return c
}

func flatten(s sexpr) (ss []sexpr) {
	_, ok := s.(cons)
	for ok {
		ss = append(ss, s.(cons).car)
		s = s.(cons).cdr
		_, ok = s.(cons)
	}
	if s != nil {
		panic("List isn't flat")
	}
	return
}
