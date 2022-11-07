FROM node:alpine3.10 as builder

WORKDIR /build
COPY . .
RUN npm ci
RUN npm run build

FROM nginx:1.23.2-alpine
WORKDIR /usr/share/nginx/html
COPY --from=builder /build/.next .next