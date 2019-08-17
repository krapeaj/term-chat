FROM golang:1.12-alpine3.9 as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/term-chat/api-gateway
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/api-gateway

FROM scratch

COPY --from=builder /go/bin/api-gateway /go/bin/api-gateway

ENTRYPOINT ["/go/bin/api-gateway"]
