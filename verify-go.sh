#!/usr/bin/env bash

cd ./lib/ff6-release4
rm -f ./ff6-release4

go build

rm -rf ./verify
mkdir ./verify
cp ../../tmp/data ./verify

./ff6-release4 tojson -i ./verify/data -o ./verify/verify.json
./ff6-release4 fromjson -i ./verify/verify.json -o ./verify/verify.sav

diff ./verify/data ./verify/verify.sav
