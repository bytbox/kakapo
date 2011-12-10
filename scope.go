package main

type scope struct {
	data   map[string]sexpr
	parent *scope
}

func (s *scope) lookup(sym string) sexpr {
	v, ok := s.data[sym]
	if ok {
		return v
	}
	if s.parent != nil {
		return s.parent.lookup(sym)
	}
	panic("undefined")
}

func (s *scope) define(sym string, val sexpr) {
	s.data[sym] = val
}

var global *scope
