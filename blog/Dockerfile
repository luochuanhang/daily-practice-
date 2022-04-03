FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/lianxi/blog
COPY . $GOPATH/src/lianxi/blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
