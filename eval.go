package main

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
		a := e.data.(atom)
		switch a.kind {
		case _SYMBOL:
		case _NUMBER:
			fmt.Printf(" number: %f\n", a.data.(float64))
		case _STRING:
			fmt.Printf(" strlit: \"%s\"\n", a.data.(string))
		default:
			panic("Invalid kind of atom")
		}
	default:
		panic("Invalid kind of sexpr")
	}
}
