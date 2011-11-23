package main

// TODO need unicode support

import (
	"fmt"
)

func eval(c chan sexpr) {
	for e := range c {
		evalSExpr(e)
	}
}

func evalSExpr(e sexpr) {
	switch e.kind {
	case _CONS:
		fmt.Println("Cons")
	case _ATOM:
		println("atom")
	default:
		panic("Invalid kind of sexpr")
	}
}
