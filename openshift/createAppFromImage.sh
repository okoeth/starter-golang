#!/bin/sh

if [ x = x$1 -o x = x$2 -o x = x$3 ]; then
    echo "Usage: ./createApp.sh "
    echo "           <namespace>"
    echo "           <app>"
    echo "           <image>"
    exit 1
fi

NAMESPACE=$1
APP=$2
IMAGE=$3

echo "Create app"
oc new-app $GIT_URL \
  --name ${APP} \
  -e MONGODB_HOST=${APP}db \
  -n $NAMESPACE \
  $IMAGE
if [ $? -ne 0 ]; then
    echo "ERROR creating app in ./createApp.sh"
    exit 1
fi

echo "Patch deploy"
./patchDeploy.sh ${APP}
if [ $? -ne 0 ]; then
    echo "ERROR patching limits in ./createApp.sh"
    exit 1
fi

echo "Create route"
oc create route edge --service=${APP}
if [ $? -ne 0 ]; then
    echo "ERROR creating route in ./createApp.sh"
    exit 1
fi

exit 0