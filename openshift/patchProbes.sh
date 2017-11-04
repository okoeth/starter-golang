#!/bin/sh
if [ x = x$1 -o x = x$2 ]; then
    echo "Usage: ./patchProbes.sh <service> <port>"
    exit
fi
SERVICE=$1
PORT=$2

# Create liveness probe
oc patch dc/$SERVICE -p '{
	"spec":{
		"template":{
			"spec":{
				"containers":[{
					"name":"'$SERVICE'",
					"readinessProbe":{
						"tcpSocket":{
							"port":'$PORT'}, 
						"initalDelaySeconds":"60",
						"timeoutSeconds":"2",
						"periodSeconds":"10",
						"successThreshold":"1",
						"failureThreshold":"3"}}]}}}}'

if [ $? -ne 0 ]; then
    echo "ERROR in ./patchProbes.sh"
    exit 1
fi

# Create liveness probe
oc patch dc/$SERVICE -p '{
	"spec":{
		"template":{
			"spec":{
				"containers":[{
					"name":"'$SERVICE'",
					"livenessProbe":{
						"tcpSocket":{
							"port":'$PORT'}, 
						"initalDelaySeconds":"60",
						"timeoutSeconds":"2",
						"periodSeconds":"10",
						"successThreshold":"1",
						"failureThreshold":"3"}}]}}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchProbes.sh"
    exit 1
fi

exit 0
