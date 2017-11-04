#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./createNamespace.sh <namespace> <app>"
    exit 1
fi
NAMESPACE=$1
APP=$2

oc get project $NAMESPACE
if [ $? -ne 0 ]; then
    echo "Creating namespace"
    oc new-project $NAMESPACE
fi

oc project $NAMESPACE

./clean.sh $NAMESPACE ${APP} 
./clean.sh $NAMESPACE ${APP}db 

exit 0
