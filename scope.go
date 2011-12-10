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

func newScope(parent *scope) *scope {
	s := new(scope)
	s.data = make(map[string]sexpr)
	s.parent = parent
	return s
}

var global *scope
