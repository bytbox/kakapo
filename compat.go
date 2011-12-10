package main

import (
	"path"
	"reflect"
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
	for name, _go := range pkg {
		global.define(pkgName+"."+name, wrapGo(_go))
	}
	return Nil
}

func wrapGo(_go interface{}) sexpr {
	typ := reflect.TypeOf(_go)
	kind := typ.Kind()
	switch kind {
	case reflect.Bool:
		b := _go.(bool)
		if b {
			return float64(1)
		} else {
			return Nil
		}
	case reflect.Int:
		return float64(_go.(int))
	case reflect.Int8:
		return float64(_go.(int8))
	case reflect.Int16:
		return float64(_go.(int16))
	case reflect.Int32:
		return float64(_go.(int32))
	case reflect.Int64:
		return float64(_go.(int64))
	case reflect.Uint:
		return float64(_go.(uint))
	case reflect.Uint8:
		return float64(_go.(uint8))
	case reflect.Uint16:
		return float64(_go.(uint16))
	case reflect.Uint32:
		return float64(_go.(uint32))
	case reflect.Uint64:
		return float64(_go.(uint64))
	case reflect.Uintptr:
		return Nil // TODO
	case reflect.Float32:
		return float64(_go.(float32))
	case reflect.Float64:
		return float64(_go.(float64))
	case reflect.Complex64:
		return Nil // TODO
	case reflect.Complex128:
		return Nil // TODO
	case reflect.Array:
		return Nil // TODO
	case reflect.Chan:
		return Nil // TODO
	case reflect.Func:
		return wrapFunc(_go)
	case reflect.Interface:
		return Nil // TODO
	case reflect.Map:
		return Nil // TODO
	case reflect.Ptr:
		return Nil // TODO
	case reflect.Slice:
		return Nil // TODO
	case reflect.String:
		return _go.(string)
	case reflect.Struct:
		return Nil // TODO
	case reflect.UnsafePointer:
		return Nil // can't handle this
	}
	return Nil
}

func wrapFunc(f interface{}) func([]sexpr) sexpr {
	// TODO patch reflect so we can do type compatibility-checking
	return func(ss []sexpr) sexpr {
		fun := reflect.ValueOf(f)
		vs := make([]reflect.Value, len(ss))
		for i, s := range ss {
			// TODO convert any cons and function arguments
			vs[i] = reflect.ValueOf(s)
		}
		r := fun.Call(vs)
		return wrapGo(r)
	}
}
