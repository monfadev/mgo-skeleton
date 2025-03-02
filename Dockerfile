# Build Stage
FROM golang:alpine AS builder

WORKDIR /go/src/mgo-skeleton

# Copy dependency files and download modules
COPY go.mod go.sum ./
RUN GIT_TERMINAL_PROMPT=0 go mod download  

# Copy source code
COPY . .  

# Build the Go application
RUN go build -v -o /go/src/mgo-skeleton/main main.go 

# Final Stage
FROM golang:alpine

RUN apk add --no-cache tzdata  

WORKDIR /go/src/mgo-skeleton

# Copy the compiled binary from the builder stage
COPY --from=builder /go/src/mgo-skeleton/main .  

# Copy .env file
COPY .env .env

# Verify the files in the container
RUN ls -l /go/src/mgo-skeleton

# Run the Go application
CMD ["./main"]
