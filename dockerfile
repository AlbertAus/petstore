# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.12.7 as builder

WORKDIR $GOPATH/src/github.com/AlbertAus/petstore
# Copy the local package files to the container's workspace.
ADD . $GOPATH/src/github.com/AlbertAus/petstore

RUN go get github.com/AlbertAus/petstore
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/AlbertAus/petstore

# Run the outyet command by default when the container starts.
CMD ["petstore"]

# Document that the service listens on port 8080.
EXPOSE 8080