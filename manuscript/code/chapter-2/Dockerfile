FROM golang:1.10-stretch

COPY src /go/src

RUN apt-get update && apt-get install -y --no-install-recommends \
        vim                                                      \
        tree