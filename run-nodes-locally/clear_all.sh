#!/usr/bin/env bash

cd ~/go/src/github.com/dispatchlabs/disgo
go build

install_files() {
  local cwd=`pwd`
  echo "Installing files in $cwd"
  rm -f -r db
  rm -f -r config
  rm -f disgo.log
  cp ~/go/src/github.com/dispatchlabs/disgo/disgo .
}

dir=~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/seed
[[ -d $dir ]] || mkdir $dir
cd $dir
install_files

for x in 1 2 3 4 5
do
  dir=~/go/src/github.com/dispatchlabs/samples/run-nodes-locally/delegate-$x
  [[ -d $dir ]] || mkdir $dir
  cd $dir
  install_files
done

