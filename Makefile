include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go parse.go eval.go util.go builtins.go math.go syntax.go compat.go packages.go
CLEANFILES = packages.go

include ${GOROOT}/src/Make.cmd

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go
	gofmt -w packages.go

scanpkgs/scanpkgs: scanpkgs/scanpkgs.${O} scanpkgs/Makefile
	${LD} -o $@ scanpkgs/scanpkgs.${O}

scanpkgs/scanpkgs.${O}: scanpkgs/main.go
	${GC} -o $@ scanpkgs/main.go

fmt:
	gofmt -w *.go
	make -C scanpkgs fmt

