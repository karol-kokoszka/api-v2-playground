FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app", "start", "--router", "chi"]