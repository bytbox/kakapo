package main

// TODO
func builtinSyntax() {

}

func builtinFunc(f func([]sexpr) sexpr) sexpr {
	return f
}

var global = scope {
	"nil": Nil,
	"+": builtinFunc(builtinAdd),
	"-": builtinFunc(builtinSub),
	"*": builtinFunc(builtinMul),
	"/": builtinFunc(builtinDiv),
}

func builtinAdd(ss []sexpr) sexpr {
	// add all numeric arguments
	r := 0.
	for _, s := range ss {
		n, ok := s.(float64)
		if !ok {
			panic("Invalid argument")
		}
		r += n
	}
	return r
}

func builtinSub(ss []sexpr) sexpr {
	if len(ss) == 0 {
		return 0
	}
	s := ss[0]
	r, ok := s.(float64)
	if !ok {
		panic("Invalid argument")
	}
	if len(ss) == 1 {
		return -r
	}
	for _, s := range ss[1:] {
		n, ok := s.(float64)
		if !ok {
			panic("Invalid argument")
		}
		r -= n
	}
	return r
}

func builtinMul(ss []sexpr) sexpr {
	// add all numeric arguments
	r := 1.
	for _, s := range ss {
		n, ok := s.(float64)
		if !ok {
			panic("Invalid argument")
		}
		r *= n
	}
	return r
}

func builtinDiv(ss []sexpr) sexpr {
	if len(ss) == 0 {
		return 0
	}
	s := ss[0]
	r, ok := s.(float64)
	if !ok {
		panic("Invalid argument")
	}
	if len(ss) == 1 {
		return 1/r
	}
	for _, s := range ss[1:] {
		n, ok := s.(float64)
		if !ok {
			panic("Invalid argument")
		}
		r /= n
	}
	return r
}
