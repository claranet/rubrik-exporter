FROM golang:1.10 

WORKDIR /go/src/github.com/claranet/rubrik-exporter
COPY . .
RUN go get
RUN go build

ENTRYPOINT [ "/go/src/github.com/claranet/rubrik-exporter/rubrik-exporter" ]

# Final image.
# FROM quay.io/prometheus/busybox:latest

# LABEL maintainer "Martin Weber <martin.weber@de.clara.net>"

# WORKDIR /
# COPY --from=builder /go/src/github.com/claranet/rubrik-exporter/rubrik-exporter /usr/local/bin/
# EXPOSE 9477
# ENTRYPOINT ["/usr/local/bin/rubrik-exporter"]
CMD [ ]