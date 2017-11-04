#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./createBuildsecretSSH.sh <namespace> <privatekeyfile>"
    exit
fi
NAMESPACE=$1
PRIVATEKEY=$2

oc secrets new-sshauth scmsecret --ssh-privatekey=$PRIVATEKEY -n $NAMESPACE
oc secrets add builder scmsecret -n $NAMESPACE
