# Stage 1: Build the Go program
FROM golang:1.20-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum 


# Copy the source code into the container
COPY . .

# Copy the .env file into the container
COPY .env .env

# Tidy the module dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o urlshortner .

# Stage 2: Create a small image for the final container
FROM alpine:latest

# Install the `curl` utility (optional, if you need to use it for any health checks)
RUN apk add --no-cache curl

# Copy the pre-built binary file from the build stage
COPY --from=build /app/urlshortner /urlshortner

# Copy the .env file from the build stage
COPY --from=build /app/.env /app/.env

# Expose the port on which the app will run (optional, if your app uses a specific port)
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["/urlshortner"]
