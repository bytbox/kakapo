package lisp

type cons struct {
	car sexpr
	cdr sexpr
}

type sexpr interface{}
type atom interface{}
type sym string
type function func(*scope, []sexpr) sexpr

type native interface{}
