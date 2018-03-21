#! /bin/bash
sudo apt-get update
sudo apt-get install -y htop
curl -O https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
tar -xzvf go1.9.4.linux-amd64.tar.gz
sudo mv go /usr/local
echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc

go get github.com/dispatchlabs/disgo
cd $GOPATH/src/github.com/dispatchlabs/disgo