#!/usr/bin/env bash

cd ./lib/ff6-release4-js

rm -rf ./verify
mkdir ./verify
cp ../../tmp/data ./verify

# node ./ff6-release4.js tojson -i ./verify/data -o ./verify/verify.json
# node ./ff6-release4.js fromjson -i ./verify/verify.json -o ./verify/verify.sav
node test.js

diff ./verify/data ./verify/verify.sav
