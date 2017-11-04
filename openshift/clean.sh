#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./clean.sh <namespace> <service>"
    exit
fi
NAMESPACE=$1
APP_NAME=$2

oc delete all,pvc -l app=$APP_NAME -n $NAMESPACE
