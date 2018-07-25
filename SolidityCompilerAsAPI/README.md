# Sample Usage
`curl -X POST 'http://35.229.93.19:9999/Compile' -d '{"codeAsBase64": "cHJhZ21hIHNvbGlkaXR5IF4wLjQuMDsKY29udHJhY3QgQmFsbG90IHsKCiAgICBzdHJ1Y3QgVm90ZXIgewogICAgICAgIHVpbnQgd2VpZ2h0OwogICAgICAgIGJvb2wgdm90ZWQ7CiAgICAgICAgdWludDggdm90ZTsKICAgICAgICBhZGRyZXNzIGRlbGVnYXRlOwogICAgfQogICAgc3RydWN0IFByb3Bvc2FsIHsKICAgICAgICB1aW50IHZvdGVDb3VudDsKICAgIH0KCiAgICBhZGRyZXNzIGNoYWlycGVyc29uOwogICAgbWFwcGluZyhhZGRyZXNzID0+IFZvdGVyKSB2b3RlcnM7CiAgICBQcm9wb3NhbFtdIHByb3Bvc2FsczsKCiAgICAvLy8gQ3JlYXRlIGEgbmV3IGJhbGxvdCB3aXRoICQoX251bVByb3Bvc2FscykgZGlmZmVyZW50IHByb3Bvc2Fscy4KICAgIGZ1bmN0aW9uIEJhbGxvdCh1aW50OCBfbnVtUHJvcG9zYWxzKSBwdWJsaWMgewogICAgICAgIGNoYWlycGVyc29uID0gbXNnLnNlbmRlcjsKICAgICAgICB2b3RlcnNbY2hhaXJwZXJzb25dLndlaWdodCA9IDE7CiAgICAgICAgcHJvcG9zYWxzLmxlbmd0aCA9IF9udW1Qcm9wb3NhbHM7CiAgICB9CgogICAgLy8vIEdpdmUgJCh0b1ZvdGVyKSB0aGUgcmlnaHQgdG8gdm90ZSBvbiB0aGlzIGJhbGxvdC4KICAgIC8vLyBNYXkgb25seSBiZSBjYWxsZWQgYnkgJChjaGFpcnBlcnNvbikuCiAgICBmdW5jdGlvbiBnaXZlUmlnaHRUb1ZvdGUoYWRkcmVzcyB0b1ZvdGVyKSBwdWJsaWMgewogICAgICAgIGlmIChtc2cuc2VuZGVyICE9IGNoYWlycGVyc29uIHx8IHZvdGVyc1t0b1ZvdGVyXS52b3RlZCkgcmV0dXJuOwogICAgICAgIHZvdGVyc1t0b1ZvdGVyXS53ZWlnaHQgPSAxOwogICAgfQoKICAgIC8vLyBEZWxlZ2F0ZSB5b3VyIHZvdGUgdG8gdGhlIHZvdGVyICQodG8pLgogICAgZnVuY3Rpb24gZGVsZWdhdGUoYWRkcmVzcyB0bykgcHVibGljIHsKICAgICAgICBWb3RlciBzdG9yYWdlIHNlbmRlciA9IHZvdGVyc1ttc2cuc2VuZGVyXTsgLy8gYXNzaWducyByZWZlcmVuY2UKICAgICAgICBpZiAoc2VuZGVyLnZvdGVkKSByZXR1cm47CiAgICAgICAgd2hpbGUgKHZvdGVyc1t0b10uZGVsZWdhdGUgIT0gYWRkcmVzcygwKSAmJiB2b3RlcnNbdG9dLmRlbGVnYXRlICE9IG1zZy5zZW5kZXIpCiAgICAgICAgICAgIHRvID0gdm90ZXJzW3RvXS5kZWxlZ2F0ZTsKICAgICAgICBpZiAodG8gPT0gbXNnLnNlbmRlcikgcmV0dXJuOwogICAgICAgIHNlbmRlci52b3RlZCA9IHRydWU7CiAgICAgICAgc2VuZGVyLmRlbGVnYXRlID0gdG87CiAgICAgICAgVm90ZXIgc3RvcmFnZSBkZWxlZ2F0ZVRvID0gdm90ZXJzW3RvXTsKICAgICAgICBpZiAoZGVsZWdhdGVUby52b3RlZCkKICAgICAgICAgICAgcHJvcG9zYWxzW2RlbGVnYXRlVG8udm90ZV0udm90ZUNvdW50ICs9IHNlbmRlci53ZWlnaHQ7CiAgICAgICAgZWxzZQogICAgICAgICAgICBkZWxlZ2F0ZVRvLndlaWdodCArPSBzZW5kZXIud2VpZ2h0OwogICAgfQoKICAgIC8vLyBHaXZlIGEgc2luZ2xlIHZvdGUgdG8gcHJvcG9zYWwgJCh0b1Byb3Bvc2FsKS4KICAgIGZ1bmN0aW9uIHZvdGUodWludDggdG9Qcm9wb3NhbCkgcHVibGljIHsKICAgICAgICBWb3RlciBzdG9yYWdlIHNlbmRlciA9IHZvdGVyc1ttc2cuc2VuZGVyXTsKICAgICAgICBpZiAoc2VuZGVyLnZvdGVkIHx8IHRvUHJvcG9zYWwgPj0gcHJvcG9zYWxzLmxlbmd0aCkgcmV0dXJuOwogICAgICAgIHNlbmRlci52b3RlZCA9IHRydWU7CiAgICAgICAgc2VuZGVyLnZvdGUgPSB0b1Byb3Bvc2FsOwogICAgICAgIHByb3Bvc2Fsc1t0b1Byb3Bvc2FsXS52b3RlQ291bnQgKz0gc2VuZGVyLndlaWdodDsKICAgIH0KCiAgICBmdW5jdGlvbiB3aW5uaW5nUHJvcG9zYWwoKSBwdWJsaWMgY29uc3RhbnQgcmV0dXJucyAodWludDggX3dpbm5pbmdQcm9wb3NhbCkgewogICAgICAgIHVpbnQyNTYgd2lubmluZ1ZvdGVDb3VudCA9IDA7CiAgICAgICAgZm9yICh1aW50OCBwcm9wID0gMDsgcHJvcCA8IHByb3Bvc2Fscy5sZW5ndGg7IHByb3ArKykKICAgICAgICAgICAgaWYgKHByb3Bvc2Fsc1twcm9wXS52b3RlQ291bnQgPiB3aW5uaW5nVm90ZUNvdW50KSB7CiAgICAgICAgICAgICAgICB3aW5uaW5nVm90ZUNvdW50ID0gcHJvcG9zYWxzW3Byb3BdLnZvdGVDb3VudDsKICAgICAgICAgICAgICAgIF93aW5uaW5nUHJvcG9zYWwgPSBwcm9wOwogICAgICAgICAgICB9CiAgICB9Cn0="}'`

# Solidity setup
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc

# Service Setup
sudo useradd dispatch-services -s /sbin/nologin -M
sudo chown -R dispatch-services:dispatch-services /go-binaries

echo '[Unit]'											| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'Description=Dispatch Compiler Service'			| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'After=network.target'								| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo '[Service]'										| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'WorkingDirectory=/go-binaries'					| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'ExecStart=/go-binaries/SolidityCompilerAsAPI'		| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'Restart=on-failure'								| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'User=dispatch-services'							| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'Group=dispatch-services'							| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo '[Install]'										| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service
echo 'WantedBy=multi-user.target'						| sudo tee --append /etc/systemd/system/dispatch-compiler-service.service

sudo systemctl enable dispatch-compiler-service
sudo systemctl start dispatch-compiler-service


# Go Setup
```
sudo apt-get update
sudo apt-get install -y htop git gcc
curl -O https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
tar -xzvf go1.9.4.linux-amd64.tar.gz
tar -xzvf go1.9.4.linux-amd64.tar.gzsudo mv go /usr/local
sudo mv go /usr/local
mkdir ~/go
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc
GOPATH=$HOME/go
PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

# Code Setup
```
go get github.com/dispatchlabs/samples
cd /home/nicu/go/src/github.com/dispatchlabs/samples/
cd SolidityCompilerAsAPI/
go build
go get ./...
go build
ls
sudo mkdir -p /go-binaries/config
sudo mv ./SolidityCompilerAsAPI /go-binaries/
```