#!/bin/sh

if [ x = x$1 ]; then
    echo "Usage: ./run_tests.sh "
    echo "           <test_server>"
    exit 1
fi

TEST_SERVER=$1

echo "STEP 000: Wait for application to become ready"
for i in 1 2 3 4 5 6 7 8 9 10; do
    curl -s -k -X GET https://$TEST_SERVER/html/
    NOT_AVAIL=`curl -s -k -X GET \
    https://$TEST_SERVER/html/ | grep "Gopher" | wc -l`
    if [ $NOT_AVAIL -eq 0 ]; then
        echo "INFO Waiting for application to become available"
        sleep 10
    fi
done

echo "STEP 001: Create record"
RECORD_ID=`curl -s -k -X POST \
  https://$TEST_SERVER/v1/greetings \
  -H 'content-type: application/json' \
  -d '{
	"title" : "Hello",
	"Message" : "Hello, Gopher!"
  }' | jq '.id' | sed -e 's/^"//' -e 's/"$//'`
if [ $? -ne 0 ]; then
    echo "ERROR creating record, curl returned code $?"
    exit 1
fi
if [ x = x$RECORD_ID ]; then
    echo "ERROR creating record, RECORD_ID is empty"
    exit 1
fi

echo "STEP 002: Retrieve record"
RECORD_TITLE=`curl -s -k -X GET \
  https://$TEST_SERVER/v1/greetings/$RECORD_ID \
  | jq '.title' | sed -e 's/^"//' -e 's/"$//'`
if [ $? -ne 0 ]; then
    echo "ERROR reading record, curl returned code $?"
    exit 1
fi
if [ xHello != x$RECORD_TITLE ]; then
    echo "ERROR reading record, RECORD_TITLE is >$RECORD_TITLE< expected >Hello<"
    exit 1
fi

echo "STEP 003: Retrieve records (minimal test, just checking return code of curl)"
curl -s -k -o /dev/null -X GET \
  https://$TEST_SERVER/v1/greetings
if [ $? -ne 0 ]; then
    echo "ERROR reading records, curl returned code $?"
    exit 1
fi

echo "STEP 004: Update record"
curl -s -k -X PUT \
  https://$TEST_SERVER/v1/greetings/$RECORD_ID \
  -H 'content-type: application/json' \
  -d '{
    "id": "'$RECORD_ID'",
	"Titel" : "Hello",
	"Message" : "THIS_CHANGED"
   }'
if [ $? -ne 0 ]; then
    echo "ERROR reading records, curl returned code $?"
    exit 1
fi
RECORD_MESSAGE=`curl -s -k -X GET \
  https://$TEST_SERVER/v1/greetings/$RECORD_ID \
  | jq '.message' | sed -e 's/^"//' -e 's/"$//'`
if [ $? -ne 0 ]; then
    echo "ERROR reading record after update, curl returned code $?"
    exit 1
fi
if [ xTHIS_CHANGED != x$RECORD_MESSAGE ]; then
    echo "ERROR reading record after update, RECORD_MESSAGE is >$RECORD_MESSAGE< expected >THIS_CHANGED<"
    exit 1
fi

echo "STEP 005: Delete record (minimal test, just checking return code of curl)"
curl -s -k -o /dev/null -X DELETE \
  https://$TEST_SERVER/v1/greetings/$RECORD_ID
if [ $? -ne 0 ]; then
    echo "ERROR deleting records, curl returned code $?"
    exit 1
fi

