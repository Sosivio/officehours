#!/bin/bash

# export CE=docker or export CE=podman
# export TAG=v1.1
#./buildall.sh

cd goapp 
go build .
chmod +x goapp
sudo $CE build -t release.sosiv.io/demoapp:$TAG .
sudo $CE push release.sosiv.io/demoapp:$TAG

cd ../be 
go build .
chmod +x be
sudo $CE build -t release.sosiv.io/demoapp-be:$TAG .
sudo $CE push release.sosiv.io/demoapp-be:$TAG