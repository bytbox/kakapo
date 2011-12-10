package main

import "fmt"

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

func (s *scope) String() string {
	return fmt.Sprintf("%s\nPARENT:\n%s", s.data, s.parent)
}

func newScope(parent *scope) *scope {
	s := new(scope)
	s.data = make(map[sym]sexpr)
	s.parent = parent
	return s
}

var global *scope
