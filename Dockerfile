FROM golang:1.22.1-bullseye as build

ENV GO11MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apt update --fix-missing && \
    apt-get upgrade -y && \
    apt install -y ca-certificates dos2unix && \
    apt install -y tzdata && \
    ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    apt-get clean

ADD go.mod go.sum ./

RUN go mod download
COPY . .

RUN dos2unix Makefile && \
    dos2unix docker-entrypoint.sh && \
    make fibo_app

FROM ubuntu:22.04 as production
RUN apt update && apt install -y ca-certificates

RUN DEBIAN_FRONTEND=noninteractive apt-get -y install tzdata

WORKDIR /app

COPY --from=build /app/fibo_app .

ENTRYPOINT ["/bin/bash", "-c"]