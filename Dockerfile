# Step 1: Build the Go app
FROM golang:1.22.6-alpine3.20 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first, then run go mod download to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -tags musl -o main

# Step 2: Run the app with a lightweight image
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Run the Go app
CMD ["./main"]
