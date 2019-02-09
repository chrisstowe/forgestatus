#!/bin/bash

gcloud beta container --project "forgestatus" clusters create "forgestatus-cluster" \
    --zone "us-west2-a" \
    --username "admin" \
    --cluster-version "1.11.6-gke.2" \
    --machine-type "n1-standard-2" \
    --image-type "COS" \
    --disk-type "pd-standard" \
    --disk-size "100" \
    --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
    --num-nodes "2" \
    --enable-cloud-logging \
    --enable-cloud-monitoring \
    --no-enable-ip-alias \
    --network "projects/forgestatus/global/networks/default" \
    --subnetwork "projects/forgestatus/regions/us-west2/subnetworks/default" \
    --addons HorizontalPodAutoscaling,HttpLoadBalancing \
    --enable-autoupgrade \
    --enable-autorepair
