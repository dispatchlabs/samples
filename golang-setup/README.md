# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_power_settings_new_black_24px.svg) Install



- Go to the download page on golang website: https://golang.org/dl/
- Choose the package file according to your operating system
- Download the package file, extract it, and follow the prompts to install the Go tools. 


##### Windows
- For Windows open the MSI file and follow the prompts to install the Go tools. By default, the installer puts the Go distribution in the c:/Go
- Ths installer should direct the Go path to the c:\Go\bin in your PATH environment variable. To check that open terminal and type `go env`
- If the GoPATH was directed to somewhere else, you can edit it through the "Environment Variables" button on the "Advaced System Settings" option inside the "System" control panel.


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
- Execute `sudo installer -pkg /your/path/go1.9.4.darwin-amd64.pkg -target /` 
then follow the promtp (you can find the path using the pwd command)
 
 __OR__
- `brew install go`



# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_build_black_24px.svg) Config

##### Windows
- On terminal you need to set up a workspace that consist of the three forlders (bin/ src/ pkg/) at the root (in this case the root is c:\projects\Go) using the following commands:
     - `mkdir c:\projects\Go`
     - `mkdir c:\projects\Go\bin`
     - `mkdir c:\projects\Go\pkg`
     - `mkdir c:\projects\Go\src`

Then set up the GOPATH

    - `set GOPATH=c:\projects\Go\bin`
    - `set PATH=%PATH%`
    
##### Dev Box
- Linux and Mac: paste the following commands in terminal:

	- `mkdir ~/go`
	- `nano ~/.bashrc`
		- `export GOPATH=$HOME/go`
		- `export PATH=$PATH:$GOPATH/bin`
		- OR
		- `export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin` - for Debian/Ubuntu



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
ExecStart=/go-binaries/disgo -asSeed -nodeId=NODE-Seed-001
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


