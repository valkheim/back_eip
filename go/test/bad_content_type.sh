curl -v -H "Content-Type: bad/one" -X POST -d '{"name":"JP","age":12}' https://127.0.0.1:443/client --cacert cacert.pem
