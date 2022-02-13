#!/bin/bash

echo '  >  Building container'
docker build -f docker/Dockerfile -t practice_k8s:v1.0 .
echo '  >  Done'
