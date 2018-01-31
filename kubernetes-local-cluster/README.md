# Concepts
Kubernetes - open source Docker orchestration.
	- handles scheduling onto nodes in a compute cluster and actively manages workloads to ensure taht their state matches the users declared intentions
	- using concepts `label` and `pod` it groups the containers which makeup an application into logical units for easy management and discovery

##### From Below
- Master Components manage nodes (aka kublets)
	- `Kublet` manages
		- `Pod` - a group of co-located group of apps running with a shared context
			- containers
			- images
			- volumes
			- etc
- When launching docker containers you do `port mapping` aka `docker run -p 80:8080`
- Kubernetes model is different: a `unique IP address` should be assigned to each `pod`
	- 1 `pod` == logically grouped docker containers
- Because pods can be on different machines, there needs to be a `network overlay` that 
ensures there can be communication between the pods on different nodes
- This where `flannel` comes in: it will give each pod an IP address they can use for 
communication to other pods between nodes
	- `flannel` uses `etcd` to store the network configuration


![](karch.png "")
![](farch.png "")

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_done_black_24px.svg) Docker Install
- Mac (https://docs.docker.com/docker-for-mac)
	- `curl -O https://download.docker.com/mac/stable/Docker.dmg`
	- Open and install the DMG file, it has all the requirements (Engine and Compose)

- Arch (https://wiki.archlinux.org/index.php/Docker)
	- `sudo pacman -S docker docker-compose`
	- `sudo systemctl enable docker`
	- `sudo systemctl start docker`
	- `sudo gpasswd -a USER docker`

- Ubuntu (https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu) (aka AWS)
	- `sudo apt-get update`
	- `sudo apt-get install apt-transport-https ca-certificates curl software-properties-common`
	- `curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -`
	- `sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"`
	- `sudo apt-get update`
	- `sudo apt-get install docker-ce`
		- `apt-cache madison docker-ce` (in PROD)
		- `sudo apt-get install docker-ce=<VERSION>` (in PROD)

##### Dox
- `docker container ls` -> list conatiners
	- `docker container stop CONTAINER_ID`
	- `docker container rm CONTAINER_ID` -> delete container, must be stopped first
- `docker images` -> list images
	- `docker rmi IMAGEID` -> delete image, add `-f` to force

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_done_black_24px.svg) Kubernetes Install

##### kubectl
- Mac (https://docs.docker.com/docker-for-mac)
	- `curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/darwin/amd64/kubectl`
	- `chmod +x ./kubectl`
	- `sudo mv ./kubectl /usr/local/bin/kubectl`

- Linux
	- `curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl`
	- `chmod +x ./kubectl`
	- `sudo mv ./kubectl /usr/local/bin/kubectl`

##### minikube
- Mac: `brew cask install minikube`
- Linux
	- `curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64`
	- `chmod +x minikube && sudo mv minikube /usr/local/bin/`

##### VirtualBox
- Download and install on your machine

##### Dox
- `kubectl get deployments`
	- `kubectl delete deployment ...`
- `kubectl get pods`
	- `kubectl delete pod ...`
- `minikube dashboard`

# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_directions_run_black_24px.svg) Run Local Cluster
- Start `Kubernetes` AND Run `Docker Private Registry` inside it
	- `minikube start` _starts a local kubernetes cluster, use `minikube delete` to start from scratch_
	- `eval $(minikube docker-env)` _points docker to the docker stuff inside the kubernetes_
	- `docker run --detach --publish 5000:5000 --name LocalPrivateRegistry registry:2` _runs a docker private registry_

- Prepare Image to run many `disgo` pods
	- `docker build -t disgo:v1 .` _-> will say `Successfully built XXXXXXXXXXX` the XXX part is `JUST_CREATED_IMAGE_ID` used below_
	- `docker tag JUST_CREATED_IMAGE_ID localhost:5000/disgo:v1`
	- `docker push localhost:5000/disgo:v1`
- Run few PODs, each will get its own IP
	- `kubectl run disgo-node1 --image=localhost:5000/disgo:v1 --port=1975 --image-pull-policy=Never`
	- `kubectl run disgo-node2 --image=localhost:5000/disgo:v1 --port=1975 --image-pull-policy=Never`
	- `kubectl run disgo-node3 --image=localhost:5000/disgo:v1 --port=1975 --image-pull-policy=Never`
	- `kubectl run disgo-node4 --image=localhost:5000/disgo:v1 --port=1975 --image-pull-policy=Never`

- Fix This
	- `kubectl create -f disgo.yaml`

![](ipport.png "")
![](minikube-dashboard.png "")