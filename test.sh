#!/bin/sh

set -e

echo "Testing..."

for f in test/*.lsp; do
	echo "  $f"
	./kakapo testing.lsp $f
done

echo "PASS" 
