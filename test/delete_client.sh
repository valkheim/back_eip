#!/bin/sh
curl --cacert cacert.pem -v -H "Content-Type: application/json" -X DELETE -d '{"id":1,"name":"JP","age":12}' https://127.0.0.1:443/client/23
