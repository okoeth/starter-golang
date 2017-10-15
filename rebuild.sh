#!/bin/bash
if [ x = x$1 ]; then
    echo "Usage: ./rebuild.sh <service>"
    echo "    <service>  : The service, e.g. starter"
  exit
fi
SERVICE=$1

docker-compose stop $SERVICE
docker-compose rm -f $SERVICE
docker-compose build $SERVICE
docker-compose create $SERVICE
docker-compose start $SERVICE
docker-compose logs $SERVICE
