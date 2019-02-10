#!/bin/bash

gcloud compute --project=forgestatus addresses create dev-forgestatus \
    --description="development address" \
    --global \
    --network-tier=PREMIUM
