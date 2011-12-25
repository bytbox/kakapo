include ${GOROOT}/src/Make.inc

TARG = lisp
BUILTINGO = cons.go math.go panic.go bool.go
SRCFILES = parse.go eval.go util.go builtins.go syntax.go compat.go primitives.go scope.go interpret.go ${BUILTINGO}
GOFILES = ${SRCFILES} init.go
#CLEANFILES = init.go
TXT2GO = ./txt2go.sh

include ${GOROOT}/src/Make.pkg

init.go: init.lsp ${TXT2GO}
	${TXT2GO} lisp init_lsp < init.lsp > $@

fmt:
	gofmt -w ${SRCFILES}

