FROM golang:1.20.5 AS builder
COPY . /src
WORKDIR /src
RUN go build -o ./bin/server

FROM busybox
COPY --from=builder /src/bin/server server
RUN chmod +x server
EXPOSE 8000