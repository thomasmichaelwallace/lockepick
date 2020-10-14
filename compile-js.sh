#!/usr/bin/env bash

rm -rf $GOPATH/src/github.com/quarnster/ff6-release4
cp -r ./lib/ff6-release4 $GOPATH/src/github.com/quarnster

cd ./lib/ff6-release4-js

$GOPATH/bin/gopherjs build github.com/quarnster/ff6-release4