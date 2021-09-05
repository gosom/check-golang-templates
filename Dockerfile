from golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go/bin/check-golang-templates

FROM alpine:latest  
WORKDIR /bin
COPY --from=builder /go/bin/check-golang-templates .
