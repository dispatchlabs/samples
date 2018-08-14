#! /bin/bash

# INPUT
# $1 - code branch
# $2 - IP:PORT where scandis must point

# Install system tools
sudo apt-get update
sudo apt-get install -y htop git gcc g++ make

# Install `nodejs` and `angular`
curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo npm install -g @angular/cli --unsafe-perm=true

# Install `nginx`
sudo apt-get install -y nginx

# Get the code
git clone -b $1 https://github.com/dispatchlabs/scandis.git
cd scandis
npm install

# Adjust the config
cp src/environments/environment.template.ts src/environments/environment.stage.ts
sed -i "s/IP:PORT/$2/g" src/environments/environment.stage.ts
ng build -c stage --aot

sudo rm -rf /var/www/html/*
sudo mv dist/scandis/* /var/www/html/