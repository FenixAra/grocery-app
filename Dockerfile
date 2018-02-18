# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9.2

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/FenixAra/grocery-app

# Build the Sessions command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/FenixAra/grocery-app

# Run the Sessions command by default when the container starts.
ENTRYPOINT /go/bin/grocery-app

# Document that the service listens on port 3000.
EXPOSE 3000
