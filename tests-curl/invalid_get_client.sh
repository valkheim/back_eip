#!/bin/sh
echo "GET on /client is not defined and should return 405 Method not allowed"
curl --cacert cacert.pem -vvv --ssl-reqd -X GET https://127.0.0.1:3000/client
