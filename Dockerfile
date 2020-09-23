FROM golang:1.15-alpine
ADD . /go/src/on-boarding
RUN go install on-boarding

FROM alpine:latest
COPY --from=0 /go/bin/on-boarding .
ENV PORT 8080
CMD ["./on-boarding"]