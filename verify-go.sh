#!/usr/bin/env bash

set -e

cd ./lib/ff6-release4
rm -f ./ff6-release4

go build

rm -rf ./verify
mkdir ./verify
cp ../../tmp/data ./verify

echo to data...
./ff6-release4 tojson -i ./verify/data -o ./verify/verify.json
echo from data...
./ff6-release4 fromjson -i ./verify/verify.json -o ./verify/verify.sav

diff ./verify/data ./verify/verify.sav
