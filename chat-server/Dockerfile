FROM golang:1.12-alpine3.9 as builder

RUN apk update && apk add --no-cache git

RUN export GOPATH=/go
WORKDIR $GOPATH/src/chat-server
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/chat-server

FROM scratch

COPY --from=builder /go/bin/chat-server /go/bin/chat-server

ENTRYPOINT ["/go/bin/chat-server"]
