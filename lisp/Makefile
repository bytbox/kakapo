include ${GOROOT}/src/Make.inc

TARG = lisp
BUILTINGO = cons.go math.go
SRCFILES = parse.go eval.go util.go builtins.go syntax.go compat.go primitives.go scope.go interpret.go ${BUILTINGO} 
GOFILES = ${SRCFILES} packages.go
CLEANFILES = packages.go

include ${GOROOT}/src/Make.pkg

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go
	gofmt -w packages.go

scanpkgs/scanpkgs: scanpkgs/scanpkgs.${O}
	${LD} -o $@ scanpkgs/scanpkgs.${O}

scanpkgs/scanpkgs.${O}: scanpkgs/main.go
	${GC} -o $@ scanpkgs/main.go

fmt:
	gofmt -w ${SRCFILES}

