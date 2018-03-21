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

echo '[Unit]'							| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Description=Dispatch Disgo Node'	| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'After=network.target'				| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo '[Service]'						| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'WorkingDirectory=/go-binaries'	| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'ExecStart=/go-binaries/disgo'		| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Restart=on-failure'				| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'User=dispatch-services'			| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'Group=dispatch-services'			| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo '[Install]'						| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service
echo 'WantedBy=multi-user.target'		| sudo tee --append /etc/systemd/system/dispatch-disgo-node.service

sudo systemctl enable dispatch-disgo-node
sudo systemctl start dispatch-disgo-node
