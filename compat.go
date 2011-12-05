package main

import (
	"path"
)

func builtinImport(ss []sexpr) sexpr {
	if len(ss) != 1 {
		panic("Invalid number of arguments")
	}

	pkgPath, ok := ss[0].(string)
	if !ok {
		panic("Invalid argument")
	}

	pkgName := path.Base(pkgPath)

	// find the package in _go_imports
	pkg, found := _go_imports[pkgPath]
	if !found {
		panic("Package not found")
	}

	// import each item
	for name, _ := range pkg {
		global[pkgName + "." + name] = Nil
	}
	return Nil
}
