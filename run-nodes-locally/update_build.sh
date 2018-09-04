#!/usr/bin/env bash

cd ~/go/src/github.com/dispatchlabs/disgo
go build

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/seed-0
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-0
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-1
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-2
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-3
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .

cd ~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-4
cp ~/go/src/github.com/dispatchlabs/disgo/disgo .
