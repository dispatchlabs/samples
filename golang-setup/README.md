# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_power_settings_new_black_24px.svg) Install

##### Debian / Ubunbu
- `sudo apt-get update`
- `sudo apt-get upgrade`
- `curl -O https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz`
- `tar -xzvf go1.9.4.linux-amd64.tar.gz`
- `sudo mv go /usr/local`

##### Arch
- `sudo pacman -S go go-tools`

##### Mac
- `curl -O https://dl.google.com/go/go1.9.4.darwin-amd64.pkg`
- Execute `go1.9.4.darwin-amd64.pkg` then next/next/finish the thing
- __OR__
- `brew install go`

##### Windows
- Download and execute `https://dl.google.com/go/go1.9.4.windows-amd64.msi`

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_build_black_24px.svg) Config

##### Dev Box
- Linux and Mac
	- `mkdir ~/go`
	- `nano ~/.bashrc`
		- `export GOPATH=$HOME/go`
		- `export PATH=$PATH:$GOPATH/bin`
		- OR
		- `export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin` - for Debian/Ubuntu

- Windows
	- `set GOPATH=c:\Users\%USERNAME%\go`
	- `mkdir C:\GOPATH\bin`
	- `mkdir C:\GOPATH\pkg`
	- `mkdir C:\GOPATH\src`
	- `set PATH=%PATH%;%GOPATH%\bin`

#### Test It
- `go env`

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_directions_run_black_24px.svg) Run Go Program

__NOTE: as example is used Dispatch `disgo` node__

##### Dev Box
- `go get github.com/dispatchlabs/disgo`
- `cd $GOPATH/src/github.com/dispatchlabs/disgo`
- `go run main.go`

##### `systemd` Linux Service
- `go get github.com/dispatchlabs/disgo`
- `cd $GOPATH/src/github.com/dispatchlabs/disgo`
- `go build`
- `sudo mkdir /go-binaries`
- `sudo mv $GOPATH/bin/disgo /go-binaries/`
- `sudo cp -r ./properties /go-binaries/`
- `sudo nano /etc/systemd/system/dispatch-disgo-node.service`
```shell
[Unit]
Description=Dispatch Disgo Node
After=network.target

[Service]
WorkingDirectory=/go-binaries
ExecStart=/go-binaries/disgo -seed -nodeId=NODE-Seed-001
Restart=on-failure

User=dispatch-services
Group=dispatch-services

[Install]
WantedBy=multi-user.target
```
- `sudo useradd dispatch-services -s /sbin/nologin -M`
- `sudo systemctl enable dispatch-disgo-node`
- `sudo systemctl start dispatch-disgo-node`
- `sudo journalctl -f -u dispatch-disgo-node`
- `sudo systemctl daemon-reload` if you change the service


