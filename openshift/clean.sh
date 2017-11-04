#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./clean.sh <namespace> <service>"
    exit -1
fi
NAMESPACE=$1
APP_NAME=$2

oc delete all,pvc -l app=$APP_NAME -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./clean.sh"
    exit -1
fi

exit 0
