#!/bin/sh
if [ x = x$1 -o x = x$2 -o x = x$3 ]; then
    echo "Usage: ./createBuildsecret.sh <namespace> <user> <password>"
    exit
fi
NAMESPACE=$1
USER=$2
PASSWORD=$3

oc secrets new-basicauth scmsecret --username=$USER --password=$PASSWORD -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createBuildSecret.sh"
    exit 1
fi

oc secrets add builder scmsecret -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createBuildSecret.sh"
    exit 1
fi

exit 0