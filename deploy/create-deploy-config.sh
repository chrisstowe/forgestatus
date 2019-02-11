#!/bin/bash

if [ ${CIRCLE_BRANCH} = 'master' ]; then
    export ENV=prod
else
    export ENV=dev
fi

mkdir config

envsubst < ./deploy/template.server.deployment.yaml > ./config/server.deployment.yaml
envsubst < ./deploy/template.server.service.yaml > ./config/server.service.yaml

for i in {1..3}
do
    export WORKER_ID=${i}
    envsubst < ./deploy/template.worker.deployment.yaml > ./config/worker.${WORKER_ID}.deployment.yaml
    envsubst < ./deploy/template.worker.service.yaml > ./config/worker.${WORKER_ID}.service.yaml
done