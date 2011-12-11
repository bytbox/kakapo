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

func (s *scope) isDefinedHere(sy sym) bool {
	_, ok := s.data[sy]
	return ok
}

func (s *scope) isDefined(sy sym) bool {
	if s.isDefinedHere(sy) {
		return true
	} else if s.parent == nil {
		return false
	}
	return s.parent.isDefined(sy)
}

func (s *scope) define(sy sym, val sexpr) {
	s.data[sy] = val
}

func (s *scope) defineHigh(sy sym, val sexpr) {
	if s.parent == nil || s.isDefinedHere(sy) {
		s.define(sy, val)
	} else {
		s.parent.defineHigh(sy, val)
	}
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
