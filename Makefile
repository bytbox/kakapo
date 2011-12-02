include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go parse.go eval.go util.go builtins.go math.go syntax.go compat.go packages.go

include ${GOROOT}/src/Make.cmd

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go

scanpkgs/scanpkgs: scanpkgs/main.go scanpkgs/Makefile
	make -C scanpkgs

fmt:
	gofmt -w *.go
	make -C scanpkgs fmt

