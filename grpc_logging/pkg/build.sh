#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o grpc-logging

export HOST=registry.cn-shenzhen.aliyuncs.com
#export HOST=harbor.cloud2go.cn
export TAG=0.1
docker build -t clarechu/grpc-logging:0.1 .

echo $DOCKER_PASSWORD| docker login --password-stdin -u clarechu

docker push clarechu/grpc-logging:0.1

rm -rf grpc-logging