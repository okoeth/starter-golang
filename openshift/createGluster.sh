#!/bin/bash
if [ x = x$1 ]; then
    echo "Usage: ./createGluster.sh <namespace>"
    exit
fi
NAMESPACE=$1

echo "Create gluster endpoint"
oc create -f glusterEndpointSvc.yml -n $NAMESPACE
