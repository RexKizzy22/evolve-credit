FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runner for the build
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main ./app

COPY .env .

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./app/main"]