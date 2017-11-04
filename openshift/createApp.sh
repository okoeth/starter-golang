#!/bin/sh

if [ x = x$1 -o x = x$2 -o x = x$3 ]; then
    echo "Usage: ./createApp.sh "
    echo "           <namespace>"
    echo "           <app>"
    echo "           <git-url>"
    echo "           <env>"
    exit
fi

NAMESPACE=$1
APP=$2
GIT_URL=$3

if [ x = x$4 ]; then
  ENV=
else
  ENV="-$4"
fi
echo "Using environment: $ENV" 

echo "Create app"
oc new-app $GIT_URL \
  --strategy=docker \
  -e MONGODB_HOST=${APP}db$ENV \
  --name ${APP}$ENV -n $NAMESPACE
if [ $? -ne 0 ]; then
    echo "ERROR in ./createApp.sh"
    exit 1
fi

echo "Patch build"
./patchBuild.sh ${APP}$ENV
if [ $? -ne 0 ]; then
    echo "ERROR in ./createApp.sh"
    exit 1
fi

echo "Patch deploy"
./patchDeploy.sh ${APP}$ENV
if [ $? -ne 0 ]; then
    echo "ERROR in ./createApp.sh"
    exit 1
fi

echo "Start build"
oc start-build ${APP}$ENV -F
if [ $? -ne 0 ]; then
    echo "ERROR in ./createApp.sh"
    exit 1
fi

echo "Create route: TODO"

exit 0