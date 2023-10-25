# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o cs-cli

# Expose any necessary ports (if your CLI app listens on a port)
# EXPOSE 8080

# Define the command to run your application
CMD ["./cs-cli"]
