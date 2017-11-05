#!/bin/sh

# Get the cert that the server is using
openssl s_client -showcerts -servername server -connect 127.0.0.1:443 > cacert.pem
