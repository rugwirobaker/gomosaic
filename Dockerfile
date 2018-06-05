#build custom image

#base image
FROM circleci/golang:1.10.0

#USER root

RUN sudo apt-get update

#download gcloud sdk
RUN curl -OL https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-203.0.0-linux-x86_64.tar.gz

#install gcloud sdk
RUN sudo tar -xz -C /usr/bin -f google-cloud-sdk-203.0.0-linux-x86_64.tar.gz &&\
    #mv google-cloud-sdk /usr/bin/. &&\
    CLOUDSDK_CORE_DISABLE_PROMPTS=1 sudo /usr/bin/google-cloud-sdk/install.sh

RUN ["/bin/bash", "-c", "source /usr/bin/google-cloud-sdk/completion.bash.inc"]

RUN ["/bin/bash", "-c", "source /usr/bin/google-cloud-sdk/path.bash.inc"]



