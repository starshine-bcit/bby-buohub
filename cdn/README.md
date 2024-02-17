# bby-buohub-cdn

This service ingests video files to produce MPEG-DASH compliant streams and then serves them.

## Features

- Receives file upload and then puts it through a multi-stage processing pipeline
- Updates database with resulting information
- Statically serves the requisite files

## Technical Details

The cdn service receives a multipart/form-data upload on the /upload endpoint. The form needs to have the "file" and "uuid" fields. Then, it looks in the database for a pre-created row containing containing a matching uuid.

The file is first inspected with gpac, then a .mpd file and chunks are created in the staging directory, where files are served from. Finally, ffmpeg is used to generate a video screenshot which can be used as a thumbnail.

## Building

You should be able to just run `go build` or `go install` from the `src` directory. Alternately, you can strip the debug symbols by running a command like this: `go install -ldflags '-s -w'`.

## Running

The cdn service relies on two environment variables being set.

1. `DB_PASSWORD`: This should equate to the database connection information in the config.
2. `SERVER_ENV`: One of `dev`, `cloud`, or `prod`. This will select between one of the *.config.yaml configuration files. If dev is selected, the http server will be run in debug mode, otherwise it will run in release mode.

You might be tempted to put the database password in the empty database/pass string in the config, don't do it! It will be overwritten from env anyway.

## Docker

The included DockerFile utilizes a multi-stage build process in order to generate a very large ubuntu image. This is necessary as the gpac binary is specifically built against Ubunutu 22.04 (at time of writing)

## API Spec

The API spec is available [here](../shared/spec/cdnspec.yml).