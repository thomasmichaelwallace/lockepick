#!/usr/bin/env bash

set -e

cd ./lib/ff6-release4-bin
rm -rf ./ff6-release4

cp -rf ../ff6-release4 ./
cp -f ./main.go ./ff6-release4
cd ./ff6-release4

go build

rm -rf ./verify
mkdir ./verify
cp ../../../tmp/data ./verify

echo to data...
./ff6-release4 tojson -i ./verify/data -o ./verify/verify.json
echo from data...
./ff6-release4 fromjson -i ./verify/verify.json -o ./verify/verify.sav

echo diff:
diff ./verify/data ./verify/verify.sav
