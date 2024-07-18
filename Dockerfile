# Use an official Golang image as a parent image
FROM golang:1.20-alpine AS build

# Set the working directory inside the container
WORKDIR /urlshortner

# Copy the source code to the workspace
COPY . .


# Compile the Go application
RUN go build -o app .

# Use a lightweight image to run the Go application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary executable from the build stage
COPY --from=build /urlshortner/app .

# Copy the .env file from the build stage
COPY --from=build /urlshortner/.env .

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Run the Go application when the container launches
CMD ["./app"]