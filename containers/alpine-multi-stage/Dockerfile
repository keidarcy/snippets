# build stage
FROM node:12-stretch AS build
WORKDIR /build
COPY package-lock.json package.json ./
RUN npm ci
COPY . .


FROM alpine:3.10
RUN apk add --update nodejs
# add a user called node
RUN addgroup -S node && adduser -S node -G node
USER node

WORKDIR /home/node/code
COPY --from=build --chown=node:node /build .
EXPOSE 9090

CMD ["node", "index.js"]
