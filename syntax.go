package main

// TODO actual pattern-matching
type syntax func([]sexpr) sexpr

// (define keyword expression)
func builtinDefine(ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	return Nil
}
