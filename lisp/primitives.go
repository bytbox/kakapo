package lisp

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

// (for cond expr)
func primitiveFor(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	cond := ss[0]
	expr := ss[1]
	val := Nil
	cv := eval(sc, cond)
	for cv != nil {
		val = eval(sc, expr)
		cv = eval(sc, cond)
	}
	return val
}

// (lambda (arg1 ...) expr)
func primitiveLambda(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	expr := ss[1]
	evalScopeParent := newScope(sc)
	var args_ = ss[0]
	// TODO type check the args list
	return function(func(callScope *scope, ss []sexpr) sexpr {
		args := args_
		evalScope := newScope(evalScopeParent)
		// Match args with ss
		aC, ok := args.(cons)
		for args != nil {
			if len(ss) == 0 {
				panic("Invalid number of arguments")
			}
			if !ok {
				// turn ss back into a cons
				val := unflatten(ss)
				s, k := args.(sym)
				if !k {
					panic("Invalid parameter specification")
				}
				evalScope.define(s, val)
				goto done
			}
			arg := aC.car
			val := ss[0]
			s, k := arg.(sym)
			if !k {
				panic("Invalid parameter specification")
			}
			evalScope.define(s, val)

			ss = ss[1:]
			args = aC.cdr
			aC, ok = args.(cons)
		}
		if len(ss) > 0 {
			panic("Invalid number of arguments")
		}
	done:
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

// (quote expr)
func primitiveQuote(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}
	return ss[0]
}

// (begin expr1 ...)
//
// This could be implemented (in large part, at least) as an ordinary function
// taking variable arguments; however, in the interest of clarity of behaviour,
// it is not.
func primitiveBegin(sc *scope, ss []sexpr) sexpr {
	last := Nil
	for _, l := range ss {
		last = eval(sc, l)
	}
	return last
}
