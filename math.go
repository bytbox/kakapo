package main

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

func builtinMod(ss []sexpr) sexpr {
	if len(ss) != 2 {
		panic("Invalid number of arguments")
	}
	a, ok1 := ss[0].(float64)
	b, ok2 := ss[1].(float64)
	if !ok1 || !ok2 {
		panic("Invalid argument")
	}
	return int(a)%int(b) // TODO fixme to work with floats
}
