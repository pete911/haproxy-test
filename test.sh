#!/bin/sh

# nodeport is set to 30000 in helm chart and kind-config

echo " --- [start cluster] ---"
kind create cluster --config kind-config.yaml --wait 5m --name haproxy-test

if [ $(kubectl config current-context) != "kind-haproxy-test" ]
then
  echo "kubectl context is not set to kind-haproxy-test"
  exit 1
fi

docker build -t pete911/haproxy-test:dev .
kind load docker-image pete911/haproxy-test:dev --name haproxy-test

# install haproxy-test chart
helm install haproxy-test charts/haproxy-test --wait

for i in {1..10}
do
   curl localhost:30000
   echo ""
done

helm delete haproxy-test
kind delete cluster --name haproxy-test
