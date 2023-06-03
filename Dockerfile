# Use the official Golang image as the base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o hello .

# Set the command to run when the container starts
CMD ["./hello"]

