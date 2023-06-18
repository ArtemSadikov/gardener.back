FROM golang:1.20-alpine AS gardener_dependencies

# Add the module files and download dependencies.
ENV GO111MODULE=on

COPY ./go.mod $GOPATH/src/gardener/go.mod
COPY ./go.sum $GOPATH/src/gardener/go.sum

WORKDIR $GOPATH/src/gardener

COPY pkg $GOPATH/src/gardener/pkg

RUN go mod download
