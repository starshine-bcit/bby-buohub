FROM node:10

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY src ./

COPY views ./

EXPOSE 8999

CMD [ "node", "start"]