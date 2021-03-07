FROM golang:alpine AS builder
WORKDIR /app
COPY app.go .
RUN go mod init app.go && go mod tidy && go build app.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
