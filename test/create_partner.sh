#!/bin/sh
curl -v -H "Content-Type: application/json" -X POST -d '{"id": 666, "name":"H&M"}' https://127.0.0.1:443/partner --cacert cacert.pem
