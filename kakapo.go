package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()

	var tc = make(chan token)
	var sc = make(chan sexpr)

	go tokenize(os.Stdin,tc)
	go parse(tc, sc)
	eval(sc)

	//<-make(chan interface{})
}
