#! /bin/bash
sudo apt-get update
sudo apt-get install -y htop git gcc

# Install GO
curl -O https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
tar -xzvf go1.9.4.linux-amd64.tar.gz
sudo mv go /usr/local
mkdir ~/go
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc