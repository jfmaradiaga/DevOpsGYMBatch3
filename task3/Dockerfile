# Use the official Go image as the base image
FROM golang:1.20.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o hello .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./hello"]