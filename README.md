# NEO4BACON

- [Documentation](https://github.com/spike-force-1-bacon-evaluators/documentation/blob/master/README.md)

- [Ways of Working](https://github.com/spike-force-1-bacon-evaluators/documentation/blob/master/docs/ways-of-working.md)


This repository contains the client and queries implementation for requesting data from Neo4j storage.

It exposes a Protocol Buffers (gRPC) for receiving requests from the BACON REST API and retrives data from storage using the Neo4j API.

## HOWTO
```
------------------------------------------------------------------------
NEO4BACON
------------------------------------------------------------------------
build/linux                    build linux binary
build/osx                      build osx binary
grpc                           generate gRPC code
image                          build docker image and push image to registry
prep                           download dependencies and format code
run                            run application container
test                           run unit tests
```
