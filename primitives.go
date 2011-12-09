package main

func primitiveIf(ss []sexpr) sexpr {
	if len(ss) < 2 || len(ss) > 3 {
		panic("Invalid number of arguments to primitive if")
	}
	cond := ss[0]
	cv := eval(cond)
	if cv != nil {
		return eval(ss[1])
	} else if len(ss) == 3 {
		return eval(ss[2])
	}
	return Nil
}

func primitiveLambda(ss []sexpr) sexpr {
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

func primitiveLet(ss []sexpr) sexpr {
	// TODO error checking
	bindings := flatten(ss[0])
	for _, _ = range bindings {

	}

	prog := ss[1:]
	last := Nil
	for _, l := range prog {
		last = eval(l)
	}
	return last
}
