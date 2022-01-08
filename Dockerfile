FROM golang:1.17.5 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o main .

FROM alpine:3.15.0
WORKDIR /app
COPY --from=builder /app/storage.json ./storage.json
COPY --from=builder /app/env.json ./env.json
COPY --from=builder /app/main ./main
EXPOSE 8000
CMD ["./main"]