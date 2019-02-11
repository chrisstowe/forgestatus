#!/bin/bash

gcloud config set project forgestatus
echo ${GCP_SERVICE_ACCOUNT} > gcp-service-account.json
gcloud auth activate-service-account --key-file gcp-service-account.json
gcloud container clusters get-credentials forgestatus-cluster --zone us-west2-a

kubectl apply -f ./config/server.deployment.yaml
kubectl apply -f ./config/server.service.yaml

WORKER_ID=1
kubectl apply -f ./config/worker.${WORKER_ID}.deployment.yaml
kubectl apply -f ./config/worker.${WORKER_ID}.service.yaml
