
FROM golang:1.22.0-alpine3.19 AS build-stage

WORKDIR /app

COPY src ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /cdn

FROM alpine:3.19 AS build-release-stage

LABEL version="0.1.0"

WORKDIR /app

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot

RUN apk update
RUN apk add curl

COPY --from=build-stage --chown=nonroot:nonroot --chmod=550 /cdn bin/cdn
COPY --chown=nonroot:nonroot --chmod=440 config/prod.config.yaml config/prod.config.yaml

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/app/bin/cdn"]

