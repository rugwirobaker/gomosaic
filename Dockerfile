#build custom image

#base image
FROM circleci/golang:1.10.0

USER root

RUN apt-get update

#download gcloud sdk
RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-203.0.0-linux-x86_64.tar.gz

#install gcloud sdk
RUN tar -xz -C /usr/bin -f google-cloud-sdk-203.0.0-linux-x86.tar.gz &&\
    CLOUDSDK_CORE_DISABLE_PROMPTS=1 google-cloud-sdk/install.sh &&\
    source /usr/bin/gcloud/filepath/path.bash.inc &&\
    source /usr/bin/gcloud/filepath/completion.bash.inc



