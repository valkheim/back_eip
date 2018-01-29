#!/bin/sh

CLI="redis-cli"

echo "flushall" | $CLI
echo "HMSET users:42 id 42 name toto" | $CLI
echo "HMSET users:1337 id 1337 name haxor" | $CLI
