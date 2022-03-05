FROM golang:alpine AS builder
COPY http_server.go /svr/
WORKDIR /svr
ENV CGO_ENABLED=0
RUN go build http_server.go

FROM scratch
COPY --from=builder /svr /svr
WORKDIR /svr
ENTRYPOINT ["./http_server"]
