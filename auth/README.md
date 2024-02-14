# bby-buohub-auth

This service supplies an authentication/authorization API for internal use by the other components of bby-buohub.

## Features

- Create user accounts
- Validate logins
- Issue and Validate JWTs

## Technical Details

The auth service requires a MariaDB service in order to function. Connection details can be specified in the config files. Note the `DB_PASSWORD` environment variable must be set in order to be able to connect.

The JWT are created using an ECDSA 384-bit keypair and loaded from the keys directory. The private key should be called `authjwt` and the public key should be `authjwt.pub`. They should both be in PEM format. The `util/key.go` file exposes the functions necessary to randomly generate a new key, if needed. Each access token is good for 15 minutes, while the refresh token is good for 24 hours.

Users are stored in a table in the database. Passwords are salted and hashed via the Argon2id algorithm. See the source in `service/cryptopts.go` for the exact parameters used.

## Building

You should be able to just run `go build` or `go install` from the `src` directory. Alternately, you can strip the debug symbols by running a command like this: `go install -ldflags '-s -w'`.

## Running

The auth service relies on two environment variables being set.

1. `DB_PASSWORD`: This should equate to the database connection information in the config.
2. `SERVER_ENV`: One of `dev`, `cloud`, or `prod`. This will select between one of the *.config.yaml configuration files. If dev is selected, the http server will be run in debug mode, otherwise it will run in release mode.

You might be tempted to put the database password in the empty database/pass string in the config, don't do it! It will be overwritten from env anyway.

## Docker

The included DockerFile utilizes a multi-stage build process in order to generate a lean container. Note that the private key is embedded in the docker image. The bby-buohub application is not really designed to deployed in the wild, so it's probably fine.

## API Spec

The API spec is available [here](../shared/spec/authspec.yml).
