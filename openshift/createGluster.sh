#!/bin/sh
if [ x = x$1 ]; then
    echo "Usage: ./createGluster.sh <namespace>"
    exit
fi
NAMESPACE=$1

echo "Create gluster endpoint"
oc create -f glusterEndpointSvc.yml -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createGluster.sh"
    exit 1
fi

exit 0
