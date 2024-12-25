ARG SERVICE_NAME=task-service
FROM golang:1.23 AS builder

ARG SERVICE_NAME
WORKDIR /app

COPY ${SERVICE_NAME}/go.mod ./${SERVICE_NAME}/go.mod
COPY ${SERVICE_NAME}/go.sum ./${SERVICE_NAME}/go.sum
COPY gen gen
RUN cd ${SERVICE_NAME} && \
  go mod download

COPY ${SERVICE_NAME}/*.go ./${SERVICE_NAME}/

RUN cd ${SERVICE_NAME} && \
  CGO_ENABLED=0 GOOS=linux go build -o app

FROM gcr.io/distroless/static-debian12

ARG SERVICE_NAME
WORKDIR /app

COPY --from=builder /app/${SERVICE_NAME}/app .

ENTRYPOINT ["/app/app"]
