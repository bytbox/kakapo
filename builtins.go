package main

// TODO
func builtinSyntax() {

}

func builtinFunc(f func([]sexpr) sexpr) sexpr {
	return sexpr{_ATOM, atom{_FUNCTION, f}}
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
		if s.kind != _ATOM {
			panic("Invalid argument")
		}
		a := s.data.(atom)
		if a.kind != _NUMBER {
			panic("Invalid argument")
		}
		n := a.data.(float64)
		r += n
	}
	return sexpr{_ATOM, atom{_NUMBER, r}}
}

func builtinSub(ss []sexpr) sexpr {
	r := 0.
	for _, s := range ss {
		if s.kind != _ATOM {
			panic("Invalid argument")
		}
		a := s.data.(atom)
		if a.kind != _NUMBER {
			panic("Invalid argument")
		}
		n := a.data.(float64)
		r += n
	}
	return sexpr{_ATOM, atom{_NUMBER, r}}
}

func builtinMul(ss []sexpr) sexpr {
	// add all numeric arguments
	r := 1.
	for _, s := range ss {
		if s.kind != _ATOM {
			panic("Invalid argument")
		}
		a := s.data.(atom)
		if a.kind != _NUMBER {
			panic("Invalid argument")
		}
		n := a.data.(float64)
		r *= n
	}
	return sexpr{_ATOM, atom{_NUMBER, r}}
}

func builtinDiv(ss []sexpr) sexpr {
	// add all numeric arguments
	r := 0.
	for _, s := range ss {
		if s.kind != _ATOM {
			panic("Invalid argument")
		}
		a := s.data.(atom)
		if a.kind != _NUMBER {
			panic("Invalid argument")
		}
		n := a.data.(float64)
		r += n
	}
	return sexpr{_ATOM, atom{_NUMBER, r}}
}
