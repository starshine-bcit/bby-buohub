# bby-buohub-genkey

This is not a service, but a small utility to generate the ECDSA-384 keys necessary to use the auth service. It outputs two files, `authjwt` and `authjwt.pub`, which are PEM encoded keyfiles.

## Building

`go build -o ../bin/genkey`

## Running

`./genkey`