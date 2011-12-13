include ${GOROOT}/src/Make.inc

TARG = lisp
BUILTINGO = cons.go math.go panic.go bool.go
SRCFILES = parse.go eval.go util.go builtins.go syntax.go compat.go primitives.go scope.go interpret.go ${BUILTINGO} 
GOFILES = ${SRCFILES}

include ${GOROOT}/src/Make.pkg

fmt:
	gofmt -w ${SRCFILES}

