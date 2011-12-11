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
	if !ok && ss[0] != nil {
		panic("Invalid argument type")
	}
	expr := ss[1]
	evalScopeParent := newScope(sc)
	var args []sexpr
	if ok {
		args = flatten(ss[0].(cons))
	} else {
		args = flatten(nil)
	}
	// TODO type check the args list
	return function(func(callScope *scope, ss []sexpr) sexpr {
		if len(ss) != len(args) {
			panic("Invalid number of arguments")
		}
		evalScope := newScope(evalScopeParent)
		for i, arg := range args {
			val := ss[i]
			evalScope.define(arg.(sym), val)
		}
		return eval(evalScope, expr)
	})
}

// (let ((sym1 val1) ...) expr1 ...)
func primitiveLet(sc *scope, ss []sexpr) sexpr {
	if len(ss) < 1 {
		panic("Invalid number of arguments")
	}
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
	sc.defineHigh(idSym, val)
	return Nil
}
