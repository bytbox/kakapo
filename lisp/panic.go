package lisp

func primitiveRecover(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 3 {
		panic("Invalid number of arguments")
	}
	return eval(sc, ss[1])
}
