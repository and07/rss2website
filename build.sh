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

figlet install Get dependencies

go get github.com/gorilla/mux

go get github.com/gosimple/slug

go get github.com/mmcdole/gofeed


#figlet install rss2website
#go get gitlab.com/and07/rss2website