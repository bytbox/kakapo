package main

type primitive func([]sexpr) sexpr

// Circumvent lame initialization loop detection. An explicit init() allows
// builtinDefine et al to reference global.
func init() {
	globalData := map[string]sexpr{
		// Primitives
		"if":     primitive(primitiveIf),
		"lambda": primitive(primitiveLambda),
		"let":    primitive(primitiveLet),
		"define": primitive(primitiveDefine),

		// Syntax (syntax.go)

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

	global = &scope{globalData, nil}
}
