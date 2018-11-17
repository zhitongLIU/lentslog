FROM golang:1.10-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . .

Run go get -u github.com/kardianos/govendor
Run govendor fetch +all

# fake download from go get
RUN mkdir -p /usr/local/go/src/github.com/zhitongLIU/lentslog/
COPY . /usr/local/go/src/github.com/zhitongLIU/lentslog/

LABEL maintainer="Zhitong Liu<zhitonggm.liu@gmail.com>"

CMD go run main.go
