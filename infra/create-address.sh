#!/bin/bash

gcloud compute --project=forgestatus addresses create forgestatus \
    --description="main ip for dashboard and api" \
    --global \
    --network-tier=PREMIUM
