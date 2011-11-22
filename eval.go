package main

// TODO need unicode support

import (
	"fmt"
)

func eval(c chan sexpr) {
	for e := range c {
		fmt.Println(e)
	}
}
