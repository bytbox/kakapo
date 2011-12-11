include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go repl.go
PREREQ = lisp
CLEANFILES = _go_.${O} ${TARG} lisp.a repl.go
TXT2GO = ./txt2go.sh

${TARG}: _go_.$O
	${LD} ${LDIMPORTS} -o $@ _go_.$O

_go_.${O}: ${GOFILES} ${PREREQ}
	$(GC) $(GCFLAGS) $(GCIMPORTS) -o $@ $(GOFILES)

repl.go: repl.lsp
	${TXT2GO} repl < repl.lsp > $@

lisp:
	make -Clisp
	cp lisp/_obj/lisp.a .

clean:
	rm -f ${CLEANFILES}
	make -Clisp clean

fmt:
	gofmt -w ${GOFILES}
	make -Clisp fmt

.PHONY: lisp

