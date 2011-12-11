package lisp

func primitiveRecover(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 3 {
		panic("Invalid number of arguments")
	}

	ids := flatten(ss[0])
	expr := ss[1]
	handler := ss[2]

	var ret sexpr
	func() {
		defer func() {
			r := recover()
			if r == nil {
				return
			}
			for _, id := range ids {
				if r == id {
					ret = eval(sc, handler)
					return
				}
			}
			panic(r)
		}()
		ret = eval(sc, expr)
	}()
	return ret
}
