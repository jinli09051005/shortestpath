#!/bin/bash

kubectl create -f config
kubectl -njinlik8s-apiserver get pods -w
