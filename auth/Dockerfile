
FROM golang:1.22.0-alpine3.19 AS build-stage

WORKDIR /app

COPY src ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth

FROM alpine:3.19 AS build-release-stage

LABEL version="0.1.0"

WORKDIR /

COPY --from=build-stage /auth /auth

EXPOSE 9000

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot

USER nonroot:nonroot

ENTRYPOINT ["/auth"]

