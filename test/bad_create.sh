#!/bin/sh
head -c 42 /dev/zero | tr '\0' '#'
echo -e "\nincomplete json"
curl -vvv -H "Content-Type: application/json" -X POST -d '{"name":"TEST2"}' https://127.0.0.1:443/client --cacert cacert.pem
