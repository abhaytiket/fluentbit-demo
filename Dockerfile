# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go mod init logapp && go mod tidy && go build -o logapp .

# Create the directory where logs will be stored
RUN mkdir -p /logs

# Set the entrypoint command to run the Go application
CMD ["./logapp"]
