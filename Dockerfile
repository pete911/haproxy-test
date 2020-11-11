FROM golang:1.15-alpine AS build
RUN apk add --no-cache gcc libc-dev

WORKDIR /go/src/app
COPY main.go .
RUN go build -o /bin/haproxy-test

FROM alpine:3.12

COPY --from=build /bin/haproxy-test /usr/local/bin/haproxy-test
ENTRYPOINT ["haproxy-test"]
