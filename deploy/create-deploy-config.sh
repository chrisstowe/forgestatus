#!/bin/bash

if [ ${CIRCLE_BRANCH} = 'master' ]; then
    export ENV=prod
else
    export ENV=dev
fi

mkdir config

envsubst < ./deploy/template.server.deployment.yaml > ./config/server.deployment.yaml
envsubst < ./deploy/template.server.service.yaml > ./config/server.service.yaml
