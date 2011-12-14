include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go repl.go packages.go
PREREQ = lisp
CLEANFILES = _go_.${O} ${TARG} lisp.a repl.go packages.go
TXT2GO = ./txt2go.sh

${TARG}: _go_.$O
	${LD} ${LDIMPORTS} -o $@ _go_.$O

_go_.${O}: ${GOFILES} ${PREREQ}
	$(GC) $(GCFLAGS) $(GCIMPORTS) -o $@ $(GOFILES)

repl.go: repl.lsp
	${TXT2GO} repl < repl.lsp > $@

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go
	gofmt -w packages.go

lisp:
	make -Clisp
	cp lisp/_obj/lisp.a .

scanpkgs/scanpkgs: scanpkgs/scanpkgs.${O}
	${LD} -o $@ scanpkgs/scanpkgs.${O}

scanpkgs/scanpkgs.${O}: scanpkgs/main.go
	${GC} -o $@ scanpkgs/main.go

clean:
	rm -f ${CLEANFILES}
	rm -f scanpkgs/scanpkgs.${O} scanpkgs/scanpkgs
	make -Clisp clean

fmt:
	gofmt -w kakapo.go
	make -Clisp fmt

test: ${TARG}
	make -Clisp test
	./test.sh

.PHONY: lisp test fmt clean

