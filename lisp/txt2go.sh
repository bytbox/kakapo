#!/bin/sh

ECHO=/bin/echo

$ECHO "package $1"
$ECHO
$ECHO -n "var $2 = \`"
cat
$ECHO "\`"

