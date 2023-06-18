FROM gardener_dependencies AS builder

ARG SERVICE_NAME=users
ARG APP=api

WORKDIR $GOPATH/src/gardener/services/${SERVICE_NAME}

COPY cmd/${APP} cmd/${APP}
COPY internal   internal
COPY config     config

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/svc ./cmd/${APP}
RUN chmod +x /go/bin/svc

FROM alpine:latest AS runner

COPY --from=builder /go/bin/svc /go/bin/svc

WORKDIR /go/bin

ENTRYPOINT [ "/go/bin/svc" ]
