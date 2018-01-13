#!/bin/sh
# Test server SSL
# curl -vvv --ssl-reqd https://127.0.0.1:443 -k # ignore CA with -k

hr()
{
  head -c 42 /dev/zero | tr '\0' '#'
  echo -e "\n$1"
  head -c 42 /dev/zero | tr '\0' '#'
  echo -e "\n"
}

hr "Simple ping"
curl --cacert cacert.pem -vvv --ssl-reqd https://127.0.0.1/debugssl
hr "Simple post"
curl --cacert cacert.pem -vvv --ssl-reqd -H "Content-Type: application/json" -X POST -d '{"id":1,"name":"JP","age":12}' https://127.0.0.1:443/client
