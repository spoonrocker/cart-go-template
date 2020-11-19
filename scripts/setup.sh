#!/bin/bash

export CART_USER=cart
export CART_PASSWORD=cart
export CART_DB=dbcart

export CART_TABLE=cart
export ITEMS_TABLE=items

echo "Installing Golang"
# update repositorie
sudo apt-get update
sudo apt-get -y upgrade
# Install Golang
cd /tmp
wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz
sudo tar -xvf go1.12.4.linux-amd64.tar.gz
sudo mv go /usr/local
# Setup go environment
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
# Check Go
go version

echo "Install PostgreSQL"
# update repositorie
sudo apt-get update
sudo apt-get -y upgrade
# Install PostgreSQL using the apt packaging system
sudo apt-get install postgresql postgresql-contrib
# Check PostgreSQL
sudo -u postgres psql -c "select version();"

echo "Creating Database User"
# Create user
sudo -u postgres psql -c "create user ${CART_USER} with password '${CART_PASSWORD}'"
sudo -u postgres psql -c "select usename from pg_catalog.pg_user"

echo "Creating Database"
# Create database
sudo -u postgres psql -c "create database ${CART_DB} owner ${CART_USER}"
sudo -u postgres psql -c "select datname from pg_catalog.pg_database"
# Assign permissions
sudo -u postgres psql -c "grant all privileges on all tables in schema public to ${CART_USER}"
sudo -u postgres psql -c "grant all privileges on all sequences in schema public to ${CART_USER}"

echo "Creating Database Tables"
# create cart table
sudo -u postgres psql -U ${CART_USER} -h 127.0.0.1 -d ${CART_DB} -c "create table ${CART_TABLE}(id serial primary key)"
sudo -u postgres psql -U ${CART_USER} -h 127.0.0.1 -d ${CART_DB} -c "select * from ${CART_TABLE}"
# create items table
sudo -u postgres psql -U ${CART_USER} -h 127.0.0.1 -d ${CART_DB} -c "create table ${ITEMS_TABLE}(id serial primary key, id_cart varchar(60) not null REFERENCES cart (id), product varchar(60) not null, quantity bigint not null)"
sudo -u postgres psql -U ${CART_USER} -h 127.0.0.1 -d ${CART_DB} -c "select * from ${ITEMS_TABLE}"