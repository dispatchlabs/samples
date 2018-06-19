
export DispatchFolder=~/go/src/github.com/dispatchlabs
export Branch=dev

mkdir -p $DispatchFolder
cd $DispatchFolder

git clone -b $Branch https://github.com/dispatchlabs/commons.git
git clone -b $Branch https://github.com/dispatchlabs/dvm.git
git clone -b $Branch https://github.com/dispatchlabs/disgover.git
git clone -b $Branch https://github.com/dispatchlabs/dapos.git
git clone -b $Branch https://github.com/dispatchlabs/disgo.git

cd disgo
go get ./...
go build
