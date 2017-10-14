#!/bin/bash

if [ x = x$1 -o x = x$2 -o x = x$3 ]; then
    echo "Usage: ./createApp.sh "
    echo "           <namespace>"
    echo "           <app>"
    echo "           <git-url>"
    echo "           <env>"
    exit
fi

NAMESPACE=$1
APP=$1
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

echo "Patch build"
./patchBuild.sh ${APP}$ENV

echo "Patch deploy"
./patchDeploy.sh ${APP}$ENV

echo "Start build"
oc start-build ${APP}$ENV -F

echo "Create route"
oc create route edge --service=${APP}$ENV
