#!/bin/sh
curl -v -H "Content-Type: application/json" https://127.0.0.1:443/client/32 --ssl-reqd --cacert cacert.pem
