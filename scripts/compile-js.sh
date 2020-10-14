#!/usr/bin/env bash

rm -rf $GOPATH/src/github.com/quarnster/ff6-release4
cp -r ./lib/ff6-release4 $GOPATH/src/github.com/quarnster

cd ./lib/ff6-release4-js

GOPHERJS_GOROOT=$HOME/sdk/go1.12.16 $GOPATH/bin/gopherjs build github.com/quarnster/ff6-release4