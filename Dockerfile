# Use the official Go base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the project files
COPY . .

# Build the Go application
RUN go build -o app

# Expose the HTTP server port
EXPOSE 8080

# Set the entrypoint command to run the Go application
CMD [ "./app" ]
