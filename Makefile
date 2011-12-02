include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go parse.go eval.go util.go builtins.go math.go

include ${GOROOT}/src/Make.cmd

fmt:
	gofmt -w *.go

