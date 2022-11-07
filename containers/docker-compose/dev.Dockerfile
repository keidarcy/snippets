FROM node:12-alpine

RUN npm i -g nodemon

USER node

WORKDIR /home/node/code

COPY --chown=node:node package.json package-lock.json ./

RUN npm ci

COPY --chown=node:node . .

CMD ["nodemon", "index.js"]
