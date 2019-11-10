FROM golang:latest AS builder
ADD . /src
RUN cd /src && go build -o goservice

# final stage
FROM centos:centos7
WORKDIR /goservice
COPY --from=builder /src/goservice /goservice/
ENTRYPOINT /goservice/goservice