#! /bin/bash
sudo apt-get update
sudo apt-get install -y htop git gcc

# Install GO
curl -O https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
tar -xzvf go1.9.4.linux-amd64.tar.gz
sudo mv go /usr/local
mkdir ~/go

# Update the ENV
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc
GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Fetch DisGo
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
sudo mv ./disgo /go-binaries/
cd /go-binaries

# Setup Disgo As Service
sudo useradd dispatch-services -s /sbin/nologin -M
sudo chown -R dispatch-services:dispatch-services /go-binaries

echo '[Unit]'							| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Description=Dispatch Disgo Node'	| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'After=network.target'				| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo '[Service]'						| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'WorkingDirectory=/go-binaries'	| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'ExecStart=/go-binaries/disgo'		| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Restart=on-failure'				| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'RestartSec=5'						| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'User=dispatch-services'			| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Group=dispatch-services'			| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo '[Install]'						| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'WantedBy=multi-user.target'		| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service

sudo systemctl enable dispatch-disgo-node
sudo systemctl start dispatch-disgo-node
