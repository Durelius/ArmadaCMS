# Build stage
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init \
    --parseDependency \
    --parseInternal \
    -g main.go \
    --output docs \
    --parseDepth 3
RUN go build -o /go/bin/app

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /app/docs ./docs
ENTRYPOINT ["/app"]
EXPOSE 3254