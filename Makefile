include ${GOROOT}/src/Make.inc

TARG = kakapo
GOFILES = kakapo.go repl.go packages.go
PREREQ = lisp
CLEANFILES = _go_.${O} ${TARG} lisp.a repl.go packages.go
TXT2GO = ./txt2go.sh

${TARG}: ${GOFILES} ${PREREQ}
	go build -o kakapo -x

repl.go: repl.lisp
	${TXT2GO} repl < repl.lisp > $@

packages.go: scanpkgs/scanpkgs
	scanpkgs/scanpkgs > packages.go
	gofmt -w packages.go

lisp:
	cd lisp && go install -x

scanpkgs/scanpkgs: scanpkgs/scanpkgs.${O}
	${LD} -o $@ scanpkgs/scanpkgs.${O}

scanpkgs/scanpkgs.${O}: scanpkgs/main.go
	${GC} -o $@ scanpkgs/main.go

clean:
	rm -f ${CLEANFILES}
	rm -f scanpkgs/scanpkgs.${O} scanpkgs/scanpkgs
	cd lisp && go clean

ifeq ($(TARGDIR),)
TARGDIR:=$(QUOTED_GOBIN)
endif

install:
	mkdir -p $(TARGDIR) && cp -f $(TARG) $(TARGDIR)

fmt:
	gofmt -w kakapo.go
	make -Clisp fmt

test: ${TARG}
	make -Clisp test
	./test.sh

.PHONY: lisp test fmt clean install

