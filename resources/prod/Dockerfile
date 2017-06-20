FROM golang:1.8.3-alpine3.6

RUN apk update && \
  apk add git

WORKDIR /go/src/github.com/spike-force-1-bacon-evaluators/neo4bacon

COPY . .

COPY resources/prod/neo4j.yml resources/conf/neo4j_local.yml

# Go get package dependencies
RUN go get -d ./...

# Run go fmt
RUN go fmt ./...

# Install application
RUN go install -v ./...

EXPOSE 50051

CMD ["neo4bacon"]
