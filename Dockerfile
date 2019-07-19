FROM golang:1.12-alpine as builder

ENV APP=serverless-functions-go

RUN apk add --update --no-cache curl bash python \
  && curl https://sdk.cloud.google.com | bash
  
ENV PATH $PATH:/root/google-cloud-sdk/bin

WORKDIR /go/src/$APP
COPY . .
CMD /go/src/$APP/scripts/functions/deploy.sh
