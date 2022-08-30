#!/bin/bash

TAG=v0.0.1
REPO=williamlehman

docker build . --target prod  -t $REPO/helloworldplugin:$TAG

docker push $REPO/helloworldplugin:$TAG