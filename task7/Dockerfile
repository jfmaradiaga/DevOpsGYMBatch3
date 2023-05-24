# Set the command to run when the container starts
CMD ["./hello"]
# Use the official Golang image as the base image
FROM golang:1.20.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN go build -o hello

# Set the entry point for the container
ENTRYPOINT ["./hello"]
