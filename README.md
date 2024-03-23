# bby-buohub

bby-buohub is a simple micro-services based containerized video streaming platform which can be run via Kubernetes or Docker Compose.

## Team members

Sasha, Dennis

## Overview

bby-buohub is composed of 4 services.

The auth service is responsible for handling user login, registration, and
issues and validates JWTs. It is not meant to be publicly accessible, instead
being called by the frontend server.

The cdn service implements an internal /upload endpoint. Files are processed
through a custom multimedia pipeline which takes advantage of GPAC and ffmpeg.
Once an upload is processed, it is available as an MPEG-DASH stream through
a static file server on this service.

The up service is the unified frontend, which allows users to login, register,
upload, and stream videos. Streaming is implemented in the browser with dash.js and video.js

Finally, the db service is just a simple MariaDB instance.

## Componenets

### Docker-Compose

In order to run the project via docker-compose, some manual setup is required after cloning this repo.

1. Generate a new set of keys with genkey tool.
2. Move the newly generated `authjwt` and `authjwt.pub` into `bby-buohub/auth/keys`
3. Follow the directions in the [up readme](./up/README.md) to create a `bby-buohub/up/.env.production` file
4. Follow the directions in the [home readme](./homepage/README.md) to create a `.env` file
5. It may? be necessary to do `npm install` and `npm run dev` once in the up project to generate the typescript definitions
6. Create a `.env` in the `bby-buohub` folder with the following template. These variables are used in the `docker-compose.yml` file

```bash
DB_PASSWORD=
SERVER_ENV=prod
DB_ROOT_PASSWORD=
```

Once those steps are complete you should be able to build and run the images!

To do that, run `docker-compose build && docker-compose up -d`

### Kubernetes

In the `kube/cloud` folder, there are a number of Kubernetes config files
which allow one to deploy bby-buohub on GKE. If you wanted to run this
yourself, you will need to precreate the cluster and also setup a
staticIPand google-managed ssl certificate, in addition to pointing your
domain to that static IP. Finally, be sure to configure and build each of
the 3 service images.

### Auth

[See the readme here](./auth/README.md)

### Cdn

[See the readme here](./cdn/README.md)

### Up

[See the readme here](./up/README.md)

### Homepage

The homepage service has been merged into the up service,
presenting a unified and scalable frontend. Users can now login,
register, upload, and view content more easily.

### Genkey

[See the readme here](./genkey/README.md)

## License

Unless otherwise specified in particular services, this project is licensed under the GNU Affero General Public License (AGPLv3). You can view the full text of the license in the LICENSE file or [here](https://www.gnu.org/licenses/agpl-3.0.txt).