FROM golang:alpine AS builder

WORKDIR /build
COPY ./ ./
RUN go build -o .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/mock-coupon-api .

EXPOSE 8080

ENTRYPOINT ["/app/mock-coupon-api"]
