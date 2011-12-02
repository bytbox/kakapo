package main

var global = scope {
	"nil": Nil,

	// Arithmetic
	"+": builtinAdd,
	"-": builtinSub,
	"*": builtinMul,
	"/": builtinDiv,
	"%": builtinMod,
}
