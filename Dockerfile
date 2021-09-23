FROM golang:1.17-alpine

RUN apk add --no-cache --update curl jq bash
RUN curl -L -o /usr/local/bin/proot https://proot.gitlab.io/proot/bin/proot \
    && chmod +x /usr/local/bin/proot