FROM golang:1.15 AS builder


WORKDIR /go/src/github.com/claranet/rubrik-exporter
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

#ENTRYPOINT [ "/go/src/github.com/claranet/rubrik-exporter/rubrik-exporter" ]

#------------------------

# Final image.
FROM quay.io/prometheus/busybox:latest
#FROM alpine
LABEL maintainer "Martin Weber <martin.weber@de.clara.net>"
WORKDIR /usr/local/bin/
COPY --from=builder /go/src/github.com/claranet/rubrik-exporter/rubrik-exporter /usr/local/bin/
EXPOSE 9477
ENTRYPOINT [ "/usr/local/bin/rubrik-exporter" ]
CMD [ ]
