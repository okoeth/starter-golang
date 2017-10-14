#!/bin/bash
if [ x = x$1 ]; then
    echo "Usage: ./patchDeploy.sh <app> <mem-limit opt.>"
    exit
fi

APP_NAME=$1

if [ x = x$2 ]; then
    echo "Default limit of 256Mi"
    MEM_LIMIT=256Mi
else
    echo "Custom limit of $2"
    MEM_LIMIT=$2
fi

# Patch missing elements in deployment
oc patch dc/$APP_NAME -p '{
    "spec":{
        "strategy":{
            "type":"Recreate"}}}'
oc patch dc/$APP_NAME -p '{
    "spec":{
        "strategy":{
            "resources":{
                "limits":{
                    "memory": "'$MEM_LIMIT'"}}}}}'
oc patch dc/$APP_NAME -p '{
    "spec":{
        "template":{
            "spec":{
                "containers":[{
                    "name":"'$APP_NAME'",
                    "resources":{
                        "limits":{
                            "memory": "'$MEM_LIMIT'"}}}]}}}}'
