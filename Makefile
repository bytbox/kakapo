include ${GOROOT}/src/Make.inc

TARG = kakapo
BUILTINGO = cons.go math.go
GOFILES = kakapo.go parse.go eval.go util.go builtins.go syntax.go compat.go packages.go primitives.go scope.go ${BUILTINGO}
CLEANFILES = packages.go

include ${GOROOT}/src/Make.cmd

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go
	gofmt -w packages.go

scanpkgs/scanpkgs: scanpkgs/scanpkgs.${O}
	${LD} -o $@ scanpkgs/scanpkgs.${O}

scanpkgs/scanpkgs.${O}: scanpkgs/main.go
	${GC} -o $@ scanpkgs/main.go

fmt:
	gofmt -w ${GOFILES}

