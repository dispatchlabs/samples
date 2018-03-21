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
source ~/.bashrc

# Compile DisGo
go get github.com/dispatchlabs/disgo
cd $GOPATH/src/github.com/dispatchlabs/disgo
go build

sudo mkdir -p /go-binaries/config
sudo mv ./disgo /go-binaries/
cd /go-binaries

# Setup Disgo As Service
sudo useradd dispatch-services -s /sbin/nologin -M
sudo chown -R dispatch-services:dispatch-services /go-binaries

echo '[Unit]'							>> /etc/systemd/system/dispatch-disgo-node.service
echo 'Description=Dispatch Disgo Node'	>> /etc/systemd/system/dispatch-disgo-node.service
echo 'After=network.targetecho'			>> /etc/systemd/system/dispatch-disgo-node.service

echo '[Service]'						>> /etc/systemd/system/dispatch-disgo-node.service
echo 'WorkingDirectory=/go-binaries'	>> /etc/systemd/system/dispatch-disgo-node.service
echo 'ExecStart=/go-binaries/disgo'		>> /etc/systemd/system/dispatch-disgo-node.service
echo 'Restart=on-failureecho'
echo 'User=dispatch-services'			>> /etc/systemd/system/dispatch-disgo-node.service
echo 'Group=dispatch-servicesecho'		>> /etc/systemd/system/dispatch-disgo-node.service

echo '[Install]'						>> /etc/systemd/system/dispatch-disgo-node.service
echo 'WantedBy=multi-user.target'		>> /etc/systemd/system/dispatch-disgo-node.service


sudo systemctl enable dispatch-disgo-node
sudo systemctl start dispatch-disgo-node
