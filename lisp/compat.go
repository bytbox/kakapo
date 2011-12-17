package lisp

import (
	"path"
	"reflect"
)

// The map of available imports
var _go_imports = map[string]map[string]interface{}{}

func ExposeImport(name string, pkg map[string]interface{}) {
	_go_imports[name] = pkg
}

// Expose an identifier globally.
func ExposeGlobal(id string, x interface{}) {
	global.define(sym(id), wrapGo(x))
}

func builtinImport(sc *scope, ss []sexpr) sexpr {
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
	for name, _go := range pkg {
		sc.define(sym(pkgName+"."+name), wrapGo(_go))
	}
	return Nil
}

func wrapGo(_go interface{}) sexpr {
	return wrapGoval(reflect.ValueOf(_go))
}

func wrapGoval(r reflect.Value) sexpr {
	typ := r.Type()
	kind := typ.Kind()
	switch kind {
	case reflect.Bool:
		b := r.Bool()
		if b {
			return float64(1)
		} else {
			return Nil
		}
	case reflect.Int:
		return float64(r.Int())
	case reflect.Int8:
		return float64(r.Int())
	case reflect.Int16:
		return float64(r.Int())
	case reflect.Int32:
		return float64(r.Int())
	case reflect.Int64:
		return float64(r.Int())
	case reflect.Uint:
		return float64(r.Uint())
	case reflect.Uint8:
		return float64(r.Uint())
	case reflect.Uint16:
		return float64(r.Uint())
	case reflect.Uint32:
		return float64(r.Uint())
	case reflect.Uint64:
		return float64(r.Uint())
	case reflect.Uintptr:
		return Nil // TODO
	case reflect.Float32:
		return float64(r.Float())
	case reflect.Float64:
		return float64(r.Float())
	case reflect.Complex64:
		return Nil // TODO
	case reflect.Complex128:
		return Nil // TODO
	case reflect.Array:
		return Nil // TODO
	case reflect.Chan:
		return Nil // TODO
	case reflect.Func:
		return wrapFunc(r.Interface())
	case reflect.Interface:
		return Nil // TODO
	case reflect.Map:
		return Nil // TODO
	case reflect.Ptr:
		return Nil // TODO
	case reflect.Slice:
		return Nil // TODO
	case reflect.String:
		return r.String()
	case reflect.Struct:
		return Nil // TODO
	case reflect.UnsafePointer:
		return Nil // can't handle this
	}
	return Nil
}

func wrapFunc(f interface{}) function {
	// TODO patch reflect so we can do type compatibility-checking
	return func(sc *scope, ss []sexpr) sexpr {
		fun := reflect.ValueOf(f)

		t := fun.Type()
		ni := t.NumIn()
		if ni != len(ss) && !t.IsVariadic() {
			panic("Invalid number of arguments")
		}

		vs := make([]reflect.Value, len(ss))
		for i, s := range ss {
			// TODO convert any cons and function arguments
			vs[i] = reflect.ValueOf(s)
		}
		r := fun.Call(vs)
		if len(r) == 0 {
			return Nil
		}
		return wrapGoval(r[0])
	}
}
