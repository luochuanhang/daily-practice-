FROM golang:latest

ENV GOPROXY http://goproxy.cn,direct
WORKDIR $GOPATH/bolg 
COPY . $GOPATH/bolg
RUN go build .
EXPOSE 8000
ENTRYPOINT [".blog"]