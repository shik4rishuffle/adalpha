FROM golang:1.11.4-alpine3.8 AS builder

WORKDIR /src
COPY . .
RUN go build -o /out/server

FROM alpine:3.8
COPY --from=builder /out/server /server
ENTRYPOINT [ "/server" ]