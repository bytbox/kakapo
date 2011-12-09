package main

type primitive func([]sexpr) sexpr

var global scope

// Circumvent lame initialization loop detection. An explicit init() allows
// builtinDefine et al to reference global.
func init() {
	global = scope{
		// Primitives
		"if":     primitive(primitiveIf),
		"lambda": primitive(primitiveLambda),
		"let":    primitive(primitiveLet),

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
