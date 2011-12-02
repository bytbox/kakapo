package main

var global = scope {
	// Syntax
	"define": builtinDefine,

	// Nil
	"nil": Nil,

	// Arithmetic
	"+": builtinAdd,
	"-": builtinSub,
	"*": builtinMul,
	"/": builtinDiv,
	"%": builtinMod,
}
