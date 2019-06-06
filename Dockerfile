FROM golang:1.12 AS builder

WORKDIR /go/src/github.com/claranet/rubrik-exporter
COPY . .
RUN go get \
    && CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

FROM scratch
COPY --from=builder /go/src/github.com/claranet/rubrik-exporter/rubrik-exporter /
EXPOSE 9477
ENTRYPOINT [ "/rubrik-exporter" ]
