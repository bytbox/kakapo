#!/bin/sh

ECHO=/bin/echo

$ECHO "package main"
$ECHO
$ECHO -n "var $1 = \`"
cat
$ECHO "\`"

