#!/bin/sh
if [ x = x$1 ]; then
    echo "Usage: ./createNamespace.sh <namespace>"
    exit 1
fi
NAMESPACE=$1

echo Namespace: $NAMESPACE
oc delete project $NAMESPACE
if [ $? -ne 0 ]; then
    echo "WARNING could not delete project" $NAMESPACE
fi

sleep 5

oc new-project $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in creating project" $NAMESPACE
    exit 1
fi

exit 0
