# bby-buohub-home

This service supplies a front-end with the capability to login, register, view available videos, and stream them from a browser.

## Features

- Create account via web browser
- Login to account via web browser
- View gallery of available videos
- Stream available video

## Technical Details

The home service is a simple express app.

It interacts with auth to set cookies in the browser containing the user's access and refresh JWT. Most requests made except to /login and /register require a user to be logged in.

Once logged in, the app will render available videos for the user to choose from. They can click on them open up a page and stream the video using video.js

## Building & Running

You will need to create a single file, `.env`. This file should be stored in the project root (`./homepage`).

```bash
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
AUTH_HOST=
AUTH_PORT=
CDN_HOST=localhost
CDN_PORT=
```

To get setup, run `npm install`, then `node src/app.js` to run.

## Docker

This app uses a prebuilt docker image and simply copies over it's source files and runs npm install.