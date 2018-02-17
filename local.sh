#!/bin/sh

# This is a script to start the api and the store at 127.0.0.1
# You must register api and store as local hosts.
# $> cat /etc/hosts
# ...
# 127.0.0.1 store
# 127.0.0.1 api
# ...
# $>
# Script starts a local redis server using a custom configuration.
# Script also starts then the GroomShop api.
# Outputs are redirected to $(pwd)/*.dump files.
# store and api are background process, remember to kill'em after use or if
# you're having trouble starting one of them (binding issue).
# The following command may be useful for such task + removing dump files
# $> pkill redis-server && pkill eip && rm -f *.dump

# Charles Paulet <charles.paulet@epitech.eu>

API="./eip"
STORE="redis-server"
HOST_API="api"
HOST_STORE="store"
STORE_CONF="redis-local.conf"
STORE_DUMP="redis-local.dump"
API_DUMP="epi-local.dump"
CMD_STORE="$STORE $STORE_CONF >$STORE_DUMP 2>&1"
CMD_API="$API >$API_DUMP 2>&1"

check_tcp() {
  ping -c 1 $1 &>/dev/null || { echo "add \"127.0.0.1 $1\" to /etc/hosts" && exit 1; }
}

rm_dumps()
{
  [ -e "$1" ] && echo "remove $1" && rm -f $1
}

run_cmd()
{
  eval "$1 &" || { echo "cannot run $1" ; exit 1; }
}

wait_for_store()
{
  echo "wait for store..."
  nc -z $HOST_STORE 6379 && { sleep .5; return; } || wait_for_store
}

check_tcp $HOST_API
check_tcp $HOST_STORE

[ ! -r "$STORE_CONF" ] && echo "$(pwd)/$STORE_CONF is not readable" && exit 1;
rm_dumps $STORE_DUMP
rm_dumps $API_DUMP

echo "starting $CMD_STORE"
pkill "$CMD_STORE"
run_cmd "$CMD_STORE"
wait_for_store
echo "starting $CMD_API"
pkill "$CMD_API"
run_cmd "$CMD_API"

exit 0
