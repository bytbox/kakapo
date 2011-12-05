package main

func builtinImport(ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}

	pkgPath, ok := ss[0].(string)
	if !ok {
		panic("Invalid argument")
	}

	// find the package in _go_imports
	_, found := _go_imports[pkgPath]
	if !found {
		panic("Package not found")
	}
	return Nil
}
