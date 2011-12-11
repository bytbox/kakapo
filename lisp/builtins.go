package lisp

type function func(*scope, []sexpr) sexpr

// Circumvent lame initialization loop detection. An explicit init() allows
// builtinDefine et al to reference global.
func init() {
	globalData := map[sym]sexpr{
		// Primitives
		"if":     primitive(primitiveIf),
		"lambda": primitive(primitiveLambda),
		"let":    primitive(primitiveLet),
		"define": primitive(primitiveDefine),

		// Nil
		"nil": Nil,

		// Cons manipulation (cons.go)
		"cons": function(builtinCons),
		"car":  function(builtinCar),
		"cdr":  function(builtinCdr),

		// Arithmetic (math.go)
		"+": function(builtinAdd),
		"-": function(builtinSub),
		"*": function(builtinMul),
		"/": function(builtinDiv),
		"%": function(builtinMod),

		// Go runtime (compat.go)
		"import": function(builtinImport),
	}

	global = &scope{globalData, nil}
}
