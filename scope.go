package main

type scope struct {
	data   map[string]sexpr
	parent *scope
}

var global *scope

