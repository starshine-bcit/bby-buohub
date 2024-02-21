FROM node:latest

WORKDIR /usr/src/app

RUN addgroup --system nonroot && adduser --system nonroot --ingroup nonroot

RUN apt update && apt upgrade -y

RUN apt install -y curl dumb-init

COPY --chown=nonroot:nonroot --chmod=440 package*.json ./

RUN npm install

USER nonroot:nonroot

COPY --chown=nonroot:nonroot --chmod=550 src ./src/

COPY --chown=nonroot:nonroot --chmod=440 .env .env

COPY --chown=nonroot:nonroot --chmod=550 views ./views/

EXPOSE 8999

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["node", "./src/app.js"]