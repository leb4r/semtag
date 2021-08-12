ARG GOVERSION="1.15.6"

FROM golang:${GOVERSION}-alpine

RUN apk add --no-cache git

WORKDIR $GOPATH/src/github.com/leb4r/semtag
COPY . .

RUN go install

WORKDIR $GOPATH
ENTRYPOINT ["semtag"]
