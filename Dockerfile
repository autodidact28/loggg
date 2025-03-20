# Use Go image to build the application
FROM golang:1.24 AS builder

# Set environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Set the working directory
WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod ./
RUN go mod download || true

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o logging-service main.go

# Use Kali Linux as the base image for running
FROM kalilinux/kali-rolling

# Install necessary runtime dependencies
RUN apt-get update && apt-get install -y \
    curl \
    awscli \
    && rm -rf /var/lib/apt/lists/*

# Copy the built executable from the builder stage
COPY --from=builder /app/logging-service /app/logging-service

# Set working directory
WORKDIR /app

# Expose any ports if needed
# EXPOSE 8080

RUN mkdir -p logs

# Run the application for 20 seconds, then upload logs to S3
CMD ["sh", "-c", "timeout 20 ./logging-service | tee logs/service.log; aws s3 sync logs s3://wecredit/"]
