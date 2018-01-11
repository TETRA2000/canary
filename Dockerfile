FROM golang:1.9.2-alpine

# From docker-gc
# https://github.com/spotify/docker-gc/blob/master/Dockerfile

ENV DOCKER_VERSION 17.09.0-ce

# We get curl so that we can avoid a separate ADD to fetch the Docker binary, and then we'll remove it
RUN apk --update add bash curl \
  && cd /tmp/ \
  && curl -sSL -O https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz \
  && tar zxf docker-${DOCKER_VERSION}.tgz \
  && mkdir -p /usr/local/bin/ \
  && mv $(find -name 'docker' -type f) /usr/local/bin/ \
  && chmod +x /usr/local/bin/docker \
  && apk del curl \
  && rm -rf /tmp/* \
  && rm -rf /var/cache/apk/*

# git
RUN apk --update add git \
    && rm -rf /var/cache/apk/*

# dep
RUN go get -u github.com/golang/dep/cmd/dep

ENV APP_HOME=$GOPATH/src/github.com/tetra2000/canary

ADD . $APP_HOME
WORKDIR $APP_HOME
RUN dep ensure
RUN go build
CMD ./canary
