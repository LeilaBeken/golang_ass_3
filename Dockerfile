# Use an official Go runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any dependencies required by the Go code
RUN go mod download

# Build the Go code into an executable
RUN go build -o app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]