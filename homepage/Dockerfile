FROM node:latest

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY src ./src/

COPY .env .env

COPY views ./views/

EXPOSE 8999

#USER nonroot:nonroot

CMD [ "node", "./src/app.js"]