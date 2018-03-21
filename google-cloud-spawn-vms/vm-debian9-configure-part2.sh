#! /bin/bash

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

sudo echo '[Unit]'							>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'Description=Dispatch Disgo Node'	>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'After=network.targetecho'			>> /etc/systemd/system/dispatch-disgo-node.servicesudo 

sudo echo '[Service]'						>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'WorkingDirectory=/go-binaries'	>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'ExecStart=/go-binaries/disgo'		>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'Restart=on-failureecho'
sudo echo 'User=dispatch-services'			>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'Group=dispatch-servicesecho'		>> /etc/systemd/system/dispatch-disgo-node.servicesudo 

sudo echo '[Install]'						>> /etc/systemd/system/dispatch-disgo-node.service
sudo echo 'WantedBy=multi-user.target'		>> /etc/systemd/system/dispatch-disgo-node.service

sudo systemctl enable dispatch-disgo-node
sudo systemctl start dispatch-disgo-node
