#!/bin/sh
if [ x = x$1 ]; then
    echo "Usage: ./patchDeploy.sh <app> <mem-limit opt.>"
    exit
fi

APP_NAME=$1

if [ x = x$2 ]; then
    echo "Default limit of 128Mi"
    MEM_LIMIT=128Mi
else
    echo "Custom limit of $2"
    MEM_LIMIT=$2
fi

# Patch missing elements in deployment
oc patch dc/$APP_NAME -p '{
    "spec":{
        "strategy":{
            "type":"Recreate"}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchDeploy.sh"
    exit 1
fi

oc patch dc/$APP_NAME -p '{
    "spec":{
        "strategy":{
            "resources":{
                "limits":{
                    "memory": "'$MEM_LIMIT'"}}}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchDeploy.sh"
    exit 1
fi

oc patch dc/$APP_NAME -p '{
    "spec":{
        "template":{
            "spec":{
                "containers":[{
                    "name":"'$APP_NAME'",
                    "resources":{
                        "limits":{
                            "memory": "'$MEM_LIMIT'"}}}]}}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchDeploy.sh"
    exit 1
fi

exit 0
