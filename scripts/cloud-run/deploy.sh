#!/usr/bin/env bash

GIT_REPO="serverless-functions-go"
ORGANIZATION="NV-Cedar"
PROJECT_NAME="cedar"
DOCKER_REGISTRY="dtr.artifacts.pwc.com/pwcusnv"
instance_name="pg-gx-e-app-700458:us-central1:cedar-database"
GCP_PROJECT_ID="pg-gx-e-app-700458"
REGION="us-central1"

# GCP_PROJECT_ID=$(gcloud config get-value project 2> /dev/null)
# gcloud beta run deploy --image=gcr.io/${projectId}/sls-fns-go --set-cloudsql-instances ${instance_name} --set-env-vars ENV=production,GCP_PROJECT=${projectId} --platform managed

COMMIT7=latest
gcloud beta run deploy ${GIT_REPO} --image=${DOCKER_REGISTRY}/${PROJECT_NAME}/${GIT_REPO}:${COMMIT7} --set-cloudsql-instances ${instance_name} --set-env-vars ENV=production,GCP_PROJECT=${GCP_PROJECT_ID} --platform managed --region ${REGION}
