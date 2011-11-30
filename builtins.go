package main

// TODO
func builtinSyntax() {

}

func builtinFunc(f func([]sexpr) sexpr) sexpr {
	return sexpr{_ATOM, atom{_FUNCTION, f}}
}

var global = scope {
	"nil": Nil,
	"+": builtinFunc(builtinPlus),
}

func builtinPlus(ss []sexpr) sexpr {
	return Nil
}
