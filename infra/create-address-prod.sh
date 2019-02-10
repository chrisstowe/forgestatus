#!/bin/bash

gcloud compute --project=forgestatus addresses create prod-forgestatus \
    --description="production address" \
    --global \
    --network-tier=PREMIUM
