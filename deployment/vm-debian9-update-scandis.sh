#! /bin/bash

# INPUT
# $1 - code branch
# $2 - IP:PORT where scandis must point

# Get the code
rm -rf scandis
git clone -b $1 https://github.com/dispatchlabs/scandis.git
cd scandis
npm install

# Adjust the config
cp src/environments/environment.template.ts src/environments/environment.stage.ts
sed -i "s/IP:PORT/$2/g" src/environments/environment.stage.ts
ng build -c stage --aot

sudo rm -rf /var/www/html/*
sudo mv dist/scandis/* /var/www/html/