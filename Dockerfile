# Debian image with Go installed, and GOPATH configured as /go
FROM golang

# Copy the local package files to the container
ADD . /go/src/github.com/goadesign/gorma-cellar
WORKDIR /go/src/github.com/goadesign/gorma-cellar

# Build the project inside the container
RUN go get
RUN go build

# Container service listens on port 8080
EXPOSE 8081
