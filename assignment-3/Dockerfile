FROM golang:1.20

ENV GOPATH=/

COPY ./ ./
# Copy the Go module files
RUN go mod download


# Build the application
RUN go build -o main .


# Expose port 8080 for the Gin HTTP server
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./main"]