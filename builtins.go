package main

var global scope

func init() {
	global = scope {
		// Syntax (syntax.go)
		"define": syntax(builtinDefine),

		// Nil
		"nil": Nil,

		// Arithmetic (math.go)
		"+": builtinAdd,
		"-": builtinSub,
		"*": builtinMul,
		"/": builtinDiv,
		"%": builtinMod,

		// Go runtime (compat.go)
		"import": builtinImport,
	}
}
