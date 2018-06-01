#!/bin/bash

echo 'Update'
echo ''
apt update
echo ''
echo ''

echo 'Upgrade'
echo ''
apt upgrade
echo ''
echo ''

echo 'Figlet'
echo ''
apt install figlet
echo ''
echo ''

figlet install golang

apt install golang

echo ''

go get github.com/gorilla/mux

go get github.com/gosimple/slug

go get github.com/mmcdole/gofeed