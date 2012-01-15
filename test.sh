#!/bin/sh

set -e

echo "Testing..."

for f in test/*.lisp; do
	echo "  $f"
	./kakapo testing.lisp $f
done

echo "PASS" 
