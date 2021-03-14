#!/usr/bin/env bash

# GO > 1.16 uses modules by default
export GO111MODULE=off

rm -rf $GOPATH/src/github.com/quarnster/ff6-release4
mkdir -p $GOPATH/src/github.com/quarnster/ff6-release4
cp -r ./lib/ff6-release4 $GOPATH/src/github.com/quarnster

cd ./lib/ff6-release4-js

# note that GOPHERJS_GOROOT should be set, as per documentation
$GOPATH/bin/gopherjs build github.com/quarnster/ff6-release4