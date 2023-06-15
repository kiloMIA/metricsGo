# Build Stage
FROM golang:1.19

WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o myapp ./server/api

# Expose the application on port 8080
EXPOSE 8080

# Run the application
CMD ["./myapp"]

