package main

// TODO actual pattern-matching
type syntax func([]sexpr) sexpr

// (define keyword expression)
func builtinDefine(ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	idSym, ok := ss[0].(sym)
	if !ok {
		panic("Invalid argument")
	}
	id := string(idSym)
	val := eval(ss[1])
	// TODO scope
	global[id] = val
	return Nil
}
