#!/bin/bash



kubectl delete -f config/

nerdctl -n k8s.io rmi -f $(nerdctl -n k8s.io images | grep di | awk '{print $3}')
