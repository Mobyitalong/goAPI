FROM golang:1.22.1 as development

ARG UID=1000
ARG GID=1001
ENV CODE_DIR=/src

USER root

RUN groupadd -g $GID gophers
RUN useradd -rm -d /home/gopher -g $GID -u $UID -s /bin/bash gopher

ENV CODE_DIR=/src
WORKDIR $CODE_DIR

COPY .$CODE_DIR $CODE_DIR
RUN chown -R gopher:gophers $CODE_DIR
RUN chmod -R 775 ${CODE_DIR}

USER gopher
