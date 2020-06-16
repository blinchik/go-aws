FROM golang:1.14 AS builder

ENV SERVICE_NAME go-aws
ENV WORKDIR ${GOPATH}/${SERVICE_NAME}

WORKDIR $WORKDIR
ADD . $WORKDIR

RUN go mod init github.com/blinchik/go-aws
RUN go install github.com/blinchik/go-aws
RUN go mod vendor

# alpine build
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' main.go

###############################################################

# alpine image
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app

WORKDIR /app

COPY --from=builder /go/go-aws/main .

CMD ["/app/main"]

ENTRYPOINT ["/app/main"]
