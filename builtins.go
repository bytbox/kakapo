package main

var global = scope {
	"nil": Nil,

	// Arithmetic
	"+": builtinAdd,
	"-": builtinSub,
	"*": builtinMul,
	"/": builtinDiv,
	"%": builtinMod,

	"define": builtinDefine,
}

func builtinDefine(ss []sexpr) sexpr {
	return Nil
}
