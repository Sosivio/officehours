#!/bin/bash

# export CE=docker or export CE=podman
# export TAG=v1.1
#./buildall.sh

cd goapp 
go build .
chmod +x goapp
$CE build -t release.sosiv.io/demoapp:$TAG .
$CE push release.sosiv.io/demoapp:$TAG

cd ../be 
go build .
chmod +x be
$CE build -t release.sosiv.io/demoapp-be:$TAG .
$CE push release.sosiv.io/demoapp-be:$TAG