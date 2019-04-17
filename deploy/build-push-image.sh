#!/bin/bash          

gcloud config set project forgestatus

echo ${GCP_SERVICE_ACCOUNT} > gcp-service-account.json

gcloud auth activate-service-account --key-file gcp-service-account.json

DOCKER_NAME=gcr.io/forgestatus/${CIRCLE_PROJECT_REPONAME}-${SERVICE}

gcloud docker -- build -t ${DOCKER_NAME}:${CIRCLE_SHA1} -f ${SERVICE}.Dockerfile .

gcloud docker -- push ${DOCKER_NAME}:${CIRCLE_SHA1}
