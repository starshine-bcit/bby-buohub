FROM node:20.11-alpine AS build-stage

RUN apk update && apk upgrade

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot

USER nonroot

WORKDIR /app

COPY --chown=nonroot:nonroot package.json ./
COPY --chown=nonroot:nonroot package-lock.json ./

RUN npm ci --omit dev

COPY --chown=nonroot:nonroot . ./

RUN npm run build

FROM node:20.11-alpine as build-release-stage

LABEL version="0.1.0"

RUN apk update && apk upgrade && apk add dumb-init curl

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot

USER nonroot

WORKDIR /app

COPY --chown=nonroot:nonroot --from=build-stage /app/build /app/package.json ./

COPY --chown=nonroot:nonroot --from=build-stage /app/node_modules ./node_modules

EXPOSE 6001

ENV HOST=0.0.0.0 PORT=6001 NODE_ENV=production

CMD ["dumb-init", "node", "-r", "dotenv/config", "index.js"]
