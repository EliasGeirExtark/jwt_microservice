### Compiler ###
FROM golang:1.19-alpine AS auth_build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Install dependences
RUN go get

# Build the application
RUN go build -o ./auth ./main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/auth .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/auth"]