
FROM golang:1.22.0-bookworm AS build-stage

WORKDIR /app

COPY src ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /cdn

FROM ubuntu:22.04 AS build-release-stage

LABEL version="0.1.0"

WORKDIR /app

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot


RUN apt update && apt upgrade -y
RUN apt install -y curl wget ffmpeg liba52-0.7.4 libfaad2 libglu1-mesa libjpeg62 libmad0
RUN wget https://download.tsi.telecom-paristech.fr/gpac/new_builds/gpac_latest_head_linux64.deb
RUN dpkg -i gpac_latest_head_linux64.deb
RUN rm gpac_latest_head_linux64.deb

COPY --from=build-stage --chown=nonroot:nonroot --chmod=550 /cdn bin/cdn
COPY --chown=nonroot:nonroot --chmod=440 config/prod.config.yaml config/prod.config.yaml

RUN mkdir /app/uploads
RUN chown nonroot:nonroot /app/uploads
RUN chmod 770 /app/uploads
VOLUME /app/uploads

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/app/bin/cdn"]

