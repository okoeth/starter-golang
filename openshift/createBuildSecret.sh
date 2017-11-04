#!/bin/sh
if [ x = x$1 -o x = x$2 -o x = x$3 ]; then
    echo "Usage: ./createBuildsecret.sh <namespace> <user> <password>"
    exit
fi
NAMESPACE=$1
USER=$2
PASSWORD=$3

SECRET_EXISTS=`oc get secrets -n $NAMESPACE | grep scmsecret | wc -l`
if [ $SECRET_EXISTS -eq 1 ]; then
    echo "Removing existing secret (ignoring errors)"

    oc secrets link builder scmsecret -n $NAMESPACE
    if [ $? -ne 0 ]; then
        echo "WARNING: Ignoring error during unlinking of secret"
    fi
    
    oc delete secret scmsecret -n $NAMESPACE
    if [ $? -ne 0 ]; then
        echo "WARNING: Ignoring error during deletion of secret"
    fi
fi

oc secrets new-basicauth scmsecret --username=$USER --password=$PASSWORD -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createBuildSecret.sh"
    exit 1
fi

oc secrets link builder scmsecret -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createBuildSecret.sh"
    exit 1
fi

exit 0