FROM golang:1.12-alpine3.9 as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/term-chat/api-gateway
COPY . .

RUN mkdir /go/bin/app
RUN cp -a certs /go/bin/app/certs
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app/api-gateway

FROM alpine

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app/api-gateway"]
