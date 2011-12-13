package lisp

import "fmt"

// (recover '(id ...) expr handler)
func builtinRecover(sc *scope, ss []sexpr) sexpr {
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
					ret = apply(sc, handler, []sexpr{r})
					return
				}
			}
			panic(r)
		}()
		ret = apply(sc, expr, []sexpr{})
		fmt.Println(ret)
	}()
	return ret
}

// (panic 'id)
func builtinPanic(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}

	id := ss[0]
	panic(id)
}
