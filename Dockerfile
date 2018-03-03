FROM alpine:3.7

LABEL MAINTAINER "Aurelien PERRIER <a.perrier89@gmail.com>"
LABEL APP "tfwrapper"

ENV TERRAFORM_PATH /root/.tfversion/bin
ENV PATH "$PATH:${TERRAFORM_PATH}"

COPY ./tfwrapper /usr/bin
