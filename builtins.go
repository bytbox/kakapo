package main

// TODO
func builtinSyntax() {

}

func builtinFunc(f func([]sexpr) sexpr) sexpr {
	return f
}

var global = scope {
	"nil": Nil,

	// Arithmetic
	"+": builtinFunc(builtinAdd),
	"-": builtinFunc(builtinSub),
	"*": builtinFunc(builtinMul),
	"/": builtinFunc(builtinDiv),
	"%": builtinFunc(builtinMod),
}
