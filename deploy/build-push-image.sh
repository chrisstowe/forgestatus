#!/bin/bash          

gcloud config set project forgestatus

echo ${GCP_SERVICE_ACCOUNT} > gcp-service-account.json

gcloud auth activate-service-account --key-file gcp-service-account.json

DOCKER_NAME=gcr.io/forgestatus/${CIRCLE_PROJECT_REPONAME}-${SERVICE}

gcloud docker -- build -t ${DOCKER_NAME}:${CIRCLE_SHA1} -f ${SERVICE}.Dockerfile .

gcloud docker -- push ${DOCKER_NAME}:${CIRCLE_SHA1}

if [ ${CIRCLE_BRANCH} = 'master' ]; then
    gcloud container images add-tag ${DOCKER_NAME}:${CIRCLE_SHA1} \
        ${DOCKER_NAME}latest
else
    gcloud container images add-tag ${DOCKER_NAME}:${CIRCLE_SHA1} \
        ${DOCKER_NAME}${CIRCLE_SHA1}-pr-`echo ${CIRCLE_PULL_REQUEST} | grep -om1 '[0-9]\+$'`
fi

gcloud container images add-tag ${DOCKER_NAME}:${CIRCLE_SHA1} \
    ${DOCKER_NAME}:${CIRCLE_SHA1}-branch-`echo ${CIRCLE_BRANCH} | sed -e 's/[^a-zA-Z0-9\-]/-/g'`
