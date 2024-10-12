# syntax=docker/dockerfile:1

FROM golang:1.23-alpine

WORKDIR $GOPATH/src/github.com/nico0302/winmonk

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["winmonk"]