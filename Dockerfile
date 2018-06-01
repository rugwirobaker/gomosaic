#build custom image

#base image
FROM circleci/golang:1.10.0

USER root

RUN apt-get update

RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-203.0.0-linux-x86_64.tar.gz

RUN tar -xz -f google-cloud-sdk-203.0.0-linux-x86_64.tar.gz

RUN CLOUDSDK_CORE_DISABLE_PROMPTS=1 google-cloud-sdk/install.sh

