FROM golang:alpine
WORKDIR /init
ADD . /init
RUN cd /init && go build
ENTRYPOINT ./init
EXPOSE 8080
