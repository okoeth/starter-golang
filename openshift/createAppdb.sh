#!/bin/bash

#TODO: Make generic

if [ x = x$1 -o x = x$2  -o x = x$3 ]; then
    echo "Usage: ./createStarterdb.sh <namespace> <storage>"
	  echo "    <namspace> : The OpenShift namespace the service is deployed in"
	  echo "    <app>      : The name of the app"
	  echo "    <storage>  : Either 'aws' or 'gluster' (specifc param w/o quotes)"
    echo "    <env>      : The environment, e.g. dev or test"
  exit
fi
NAMESPACE=$1
APP=$2
STORAGE=$3

if [ x = x$4 ]; then
  ENV=
else
  ENV="-$4"
fi
echo "Using environment: $ENV" 

# Create mongodb service, deployment etc.
oc new-app \
    --name=${APP}db$ENV \
	  -n $NAMESPACE \
    -e MONGODB_USER=user \
    -e MONGODB_PASSWORD=password \
    -e MONGODB_DATABASE=${APP}db \
    -e MONGODB_ADMIN_PASSWORD=admin_password \
    centos/mongodb-26-centos7

# Create volume claim
echo "Create volume claim"
sed 's/APP/'$APP'/g' < appdb_claim_$STORAGE.yaml > ${APP}db_claim_$STORAGE.yaml
oc create -f ${APP}db_claim_$STORAGE.yaml
rm ${APP}db_claim_$STORAGE.yaml

# Attach volume claim
echo "Attach volume claim"
oc volume dc/${APP}db$ENV --add --overwrite \
  --name=${APP}db-volume-1 --type=persistentVolumeClaim \
  --claim-name=${APP}db-claim --mount-path=//var/lib/mongodb

# Adjust limits and deployment strategy
echo "Adjust limits and deployment strategy"
./patchDeploy.sh ${APP}db$ENV

# Create probes
echo "Create probes"
./patchProbes.sh ${APP}db$ENV 27017
