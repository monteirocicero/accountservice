#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0


echo 'Building services'
# accountservice
go get;go build -o accountservice-linux-amd64;echo built `pwd`;cd ..

# health check service
cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..


export GOOS=darwin

# copy healthchecker binary into accountservice
cp healthchecker/healthchecker-linux-amd64 accountservice/

echo 'Docker tagging and push to registry'

docker build -t monteirocicero/accountservice accountservice/
docker tag monteirocicero/accountservice localhost:5000/my-accountservice
docker push localhost:5000/my-accountservice
docker image remove localhost:5000/my-accountservice

echo 'Deploy on Swarm'

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=my_network -p=6767:6767 localhost:5000/my-accountservice
