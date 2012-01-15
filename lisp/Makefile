include ${GOROOT}/src/Make.inc

TARG = lisp
BUILTINGO = cons.go math.go panic.go bool.go
SRCFILES = parse.go eval.go util.go builtins.go syntax.go compat.go types.go primitives.go scope.go interpret.go ${BUILTINGO}
GOFILES = ${SRCFILES} init.go
#CLEANFILES = init.go
TXT2GO = ./txt2go.sh

include ${GOROOT}/src/Make.pkg

init.go: init.lisp ${TXT2GO}
	${TXT2GO} lisp init_lisp < init.lisp > $@

fmt:
	gofmt -w ${SRCFILES}

