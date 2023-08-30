FROM golang:1.19-alpine AS builder

RUN apk add git

WORKDIR /app
COPY cmd cmd
COPY internal internal
COPY utils utils
COPY go.mod go.mod

RUN go mod tidy;
RUN go mod vendor;
RUN go build -o service cmd/service/main.go

FROM bash:4.4
COPY --from=builder /app/service /service
COPY configuration/config.yaml configuration/config.yaml

CMD ["/service"]

