#!/bin/bash

#set -x

echo "Installing Golang"
cd /tmp
wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
tar -xvf go1.13.4.linux-amd64.tar.gz
mv go /usr/local

# setup go environment
export GOROOT=/usr/local/go
export GOPATH=/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# check go version
go version

echo "Installing postgres"
apt-get install -y  postgresql
service postgresql start

# check postgres version
sudo -u postgres psql -c "select version();"

# setup database environment
export CART_USER=cart
export CART_PASSWORD=cart
export CART_DB=dbcart
export CART_TABLE=cart
export ITEMS_TABLE=items

echo "Creating database user cart with password cart"
sudo -u postgres psql -c "create user ${CART_USER} with password '${CART_PASSWORD}';"
sudo -u postgres psql -c "select usename from pg_catalog.pg_user;"

echo "Creating database dbcart"
sudo -u postgres psql -c "create database ${CART_DB} owner ${CART_USER};"
sudo -u postgres psql -c "select datname from pg_catalog.pg_database;"

echo "Assigning permissions to cart user"
sudo -u postgres psql -c "grant all privileges on all tables in schema public to ${CART_USER};"
sudo -u postgres psql -c "grant all privileges on all sequences in schema public to ${CART_USER};"

echo "Creating database tables"
sudo PGPASSWORD=$CART_PASSWORD -u postgres psql -U ${CART_USER} -h localhost -d ${CART_DB} -c "create table ${CART_TABLE}(id serial primary key, date date not null);"
sudo PGPASSWORD=$CART_PASSWORD -u postgres psql -U ${CART_USER} -h localhost -d ${CART_DB} -c "select * from ${CART_TABLE};"

sudo PGPASSWORD=$CART_PASSWORD -u postgres psql -U ${CART_USER} -h localhost -d ${CART_DB} -c "create table ${ITEMS_TABLE}(id serial, cart_id int references ${CART_TABLE} (id), product varchar(100) not null, quantity int not null, primary key(id, cart_id));"
sudo PGPASSWORD=$CART_PASSWORD -u postgres psql -U ${CART_USER} -h localhost -d ${CART_DB} -c "select * from ${ITEMS_TABLE};"

echo "Building project"
cd /go/src/github.com/Kleiber/cart-go-template/src/
go mod download
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

