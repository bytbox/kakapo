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
