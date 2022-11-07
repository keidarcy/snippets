FROM alpine:3.10

RUN apk add --update nodejs npm

# add a user called node
RUN addgroup -S node && adduser -S node -G node

USER node

WORKDIR /home/node/code

COPY --chown=node:node package.json package-lock.json ./

RUN npm ci

COPY --chown=node:node  . .
# ADD --chown=node:node index.js index.js
# ADD let you get from network, COPY only from local
# ADD unzips the file, COPY does not


EXPOSE 9090

CMD ["node", "index.js"]
