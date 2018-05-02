export Branch=dev

# Fetch
git clone -b $Branch https://github.com/dispatchlabs/commons.git
git clone -b $Branch https://github.com/dispatchlabs/dapos.git
git clone -b $Branch https://github.com/dispatchlabs/disgover.git
git clone -b $Branch https://github.com/dispatchlabs/disgo.git

# Pull Dependencies
cd commons
go get ./...

cd ../dapos
go get ./...

cd ../disgover
go get ./...

cd ../disgo
go get ./...