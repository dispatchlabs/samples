#! /bin/bash

# Fetch DisGo
rm -rf $GOPATH/src/github.com/dispatchlabs
mkdir -p $GOPATH/src/github.com/dispatchlabs
cd $GOPATH/src/github.com/dispatchlabs
git clone -b $1 https://github.com/dispatchlabs/disgo.git

# Pull Dependencies
cd disgo
go get ./...

# Compile DisGo
cd $GOPATH/src/github.com/dispatchlabs/disgo
go build

sudo mkdir -p /go-binaries/config
sudo systemctl stop dispatch-disgo-node
sudo mv ./disgo /go-binaries/
sudo systemctl start dispatch-disgo-node
