#!/bin/bash

gcloud dns --project=forgestatus record-sets transaction start --zone=forgestatus-com

gcloud dns --project=forgestatus record-sets transaction add 130.211.8.101 --name=dev.forgestatus.com. --ttl=300 --type=A --zone=forgestatus-com

gcloud dns --project=forgestatus record-sets transaction execute --zone=forgestatus-com
