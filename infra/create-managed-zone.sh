#!/bin/bash

gcloud beta dns --project=forgestatus managed-zones create forgestatus-com \
    --description="main zone" \
    --dns-name=forgestatus.com.
