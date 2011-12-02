package main

var global = scope {
	// Syntax
	"define": syntax(builtinDefine),

	// Nil
	"nil": Nil,

	// Arithmetic
	"+": builtinAdd,
	"-": builtinSub,
	"*": builtinMul,
	"/": builtinDiv,
	"%": builtinMod,

	// Go runtime
	"import": builtinImport,
}
