# Start from the official Go image
FROM golang:latest

# Set necessary environment variables
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose any necessary ports (if applicable)
# EXPOSE 

# Command to run the CLI command (serve-api) inside the container
CMD ["./main", "serve-api"]
