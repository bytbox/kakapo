package lisp

// (not b)
func builtinNot(sc *scope, ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}
	if ss[0] == nil {
		return 1.0
	}
	return nil
}
