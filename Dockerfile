# syntax=docker/dockerfile:1

FROM golang:1.16

WORKDIR /bin

ADD ./ /bin/curd_demo/