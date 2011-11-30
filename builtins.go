package main

// TODO
func builtinSyntax() {

}

func builtinFunc(f func([]sexpr) sexpr) sexpr {
	return sexpr{_ATOM, atom{_FUNCTION, f}}
}

var global = scope {
	"nil": Nil,
	"+": builtinFunc(builtinPlus),
	"-": Nil,
	"*": Nil,
	"/": Nil,
}

func builtinPlus(ss []sexpr) sexpr {
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
