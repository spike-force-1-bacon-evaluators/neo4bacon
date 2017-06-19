FROM golang:1.8.3-alpine3.6

RUN apk update && \
  apk add git

WORKDIR /go/src/github.com/bacon/neo4bacon

COPY . .

# Go get package dependencies
RUN go get -d ./...

# Run go fmt
RUN go fmt ./...

# Run unit tests
CMD ["go", "test", "./...", "-v"]
