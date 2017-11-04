#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./createNamespace.sh <prefix> <branch>"
    exit 1
fi
PREFIX=$1
BRANCH=$2

echo Namespace: $PREFIX-$BRANCH

exit 0
