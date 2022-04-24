FROM golang:alpine AS builder
COPY httpserver /svr/
WORKDIR /svr
ENV CGO_ENABLED=0
RUN go build

#FROM scratch
FROM debian
COPY --from=builder /svr /svr
WORKDIR /svr
ENTRYPOINT ["./httpserver"]
