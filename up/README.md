# bby-buohub-auth

This service supplies a front-end with the capability to login, register and upload videos. It depends on cdn, auth, and the database being up.

## Features

- Create account via web browser
- Login to account via web browser
- Upload video files

## Technical Details

The up service is a SvelteKit app using the NodeJS adapator, which allows it to be easily packaged with docker.

It interacts with auth to set cookies in the browser containing the user's access and refresh JWT. Any requests made except to /login and /register require a user to be logged in.

When a video is uploaded, it creates an entry in the database, then forwards that video to cdn's /upload endpoint, where it is meant to automatically processed.

## Building & Running

You will need to create two files, `.env.development` and `.env.production` in the following format. Be sure to use suitable values for your environment. These files should be stored in the project root (`./up`).

```bash
HOST=
PORT=
DB_PASSWORD=
DB_PORT=
DB_USER=
DB_HOST=
DB_NAME=
AUTH_HOST=
AUTH_PORT=
CDN_HOST=
CDN_PORT=
HOME_HOST=
HOST_PORT=
```

You can get setup by running `npm install`. Then, the dev version can be run with `npm run dev`.

To build, you can run `npm run build` and then `npm -r dotenv/config build` to run it.

## Docker

The included DockerFile utilizes a multi-stage build process in order to generate a lean container. Note that passwords are embedded in the container, but this is probably fine since bby-buohub isn't meant to be actually used.
