package main

type scope struct {
	data   map[sym]sexpr
	parent *scope
}

func (s *scope) lookup(sy sym) sexpr {
	v, ok := s.data[sy]
	if ok {
		return v
	}
	if s.parent != nil {
		return s.parent.lookup(sy)
	}
	panic("undefined")
}

func (s *scope) define(sy sym, val sexpr) {
	s.data[sy] = val
}

func newScope(parent *scope) *scope {
	s := new(scope)
	s.data = make(map[sym]sexpr)
	s.parent = parent
	return s
}

var global *scope
