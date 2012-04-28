#!/bin/sh

set -e

./genrepl.sh
(cd scanpkgs && go build)
scanpkgs/scanpkgs > packages.go
go build ./lisp
go build
