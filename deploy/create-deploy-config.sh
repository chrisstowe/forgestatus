#!/bin/bash

set -e

if [ ${CIRCLE_BRANCH} = 'master' ]; then
    export ENV=prod
else
    export ENV=dev
fi

mkdir config

envsubst < ./deploy/template.namespace.yaml > ./config/namespace.yaml
envsubst < ./deploy/template.ingress.yaml > ./config/ingress.yaml

envsubst < ./deploy/template.redis.deployment.yaml > ./config/redis.deployment.yaml
envsubst < ./deploy/template.redis.service.yaml > ./config/redis.service.yaml

export WORKER_COUNT=4
for i in $( seq 1 $WORKER_COUNT )
do
    export WORKER_ID=${i}
    envsubst < ./deploy/template.worker.deployment.yaml > ./config/worker.${WORKER_ID}.deployment.yaml
    envsubst < ./deploy/template.worker.service.yaml > ./config/worker.${WORKER_ID}.service.yaml
done

envsubst < ./deploy/template.server.deployment.yaml > ./config/server.deployment.yaml
envsubst < ./deploy/template.server.service.yaml > ./config/server.service.yaml
