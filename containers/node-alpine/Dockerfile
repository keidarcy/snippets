FROM node:12-alpine

USER node

WORKDIR /home/node/code

COPY --chown=node:node package.json package-lock.json ./

RUN npm ci

COPY --chown=node:node . .
# ADD --chown=node:node index.js index.js
# ADD let you get from network, COPY only from local
# ADD unzips the file, COPY does not


EXPOSE 3000

CMD ["node", "index.js"]
