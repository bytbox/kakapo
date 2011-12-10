package main

type primitive func(*scope, []sexpr) sexpr

// (if cond expr1 expr2)
func primitiveIf(sc *scope, ss []sexpr) sexpr {
	if len(ss) < 2 || len(ss) > 3 {
		panic("Invalid number of arguments to primitive if")
	}
	cond := ss[0]
	cv := eval(sc, cond)
	if cv != nil {
		return eval(sc, ss[1])
	} else if len(ss) == 3 {
		return eval(sc, ss[2])
	}
	return Nil
}

// (lambda (arg1 ...) expr)
func primitiveLambda(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	_, ok := ss[0].(cons)
	if !ok {
		panic("Invalid argument type")
	}
	// TODO
	return Nil
}

// (let ((sym1 val1) ...) expr1 ...)
func primitiveLet(sc *scope, ss []sexpr) sexpr {
	evalScope := newScope(sc)
	bindings := flatten(ss[0])
	for _, b := range bindings {
		bs := flatten(b)
		if len(bs) != 2 {
			panic("Invalid binding")
		}
		s, ok := bs[0].(sym)
		if !ok {
			panic("Invalid binding")
		}
		val := eval(sc, bs[1])
		evalScope.define(s, val)
	}

	prog := ss[1:]
	last := Nil
	for _, l := range prog {
		last = eval(evalScope, l)
	}
	return last
}

// (define keyword expression)
func primitiveDefine(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	idSym, ok := ss[0].(sym)
	if !ok {
		panic("Invalid argument")
	}
	val := eval(sc, ss[1])
	// TODO *scope
	sc.define(idSym, val)
	return Nil
}
