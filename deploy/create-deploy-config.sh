#!/bin/bash

if [ ${CIRCLE_BRANCH} = 'master' ]; then
    export ENV=prod
else
    export ENV=dev
fi

mkdir config

envsubst < ./deploy/template.server.deployment.yaml > ./config/server.deployment.yaml
envsubst < ./deploy/template.server.service.yaml > ./config/server.service.yaml

# WORKER_ID is used in the template and the generated file name
export WORKER_ID=1
envsubst < ./deploy/template.worker.deployment.yaml > ./config/worker.${WORKER_ID}.deployment.yaml
envsubst < ./deploy/template.worker.service.yaml > ./config/worker.${WORKER_ID}.service.yaml
