#!/bin/bash

gcloud dns --project=forgestatus record-sets transaction start --zone=forgestatus-com

gcloud dns --project=forgestatus record-sets transaction add 35.190.38.85 --name=forgestatus.com. --ttl=300 --type=A --zone=forgestatus-com

gcloud dns --project=forgestatus record-sets transaction execute --zone=forgestatus-com
