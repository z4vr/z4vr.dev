# Buidling web application
FROM node:16-alpine AS web

WORKDIR /build

COPY public/ public/
COPY src/ src/
COPY package.json .
COPY package-lock.json .
COPY vue.config.js .

RUN npm ci
RUN npm run build

# Build shitsuji webserver
FROM golang:1.18-alpine as golang

WORKDIR /build

RUN apk add git gcc make
RUN git clone https://github.com/z4vr/shitsuji . --depth 1
RUN go build -o shitsuji cmd/main.go

# Finalize
FROM alpine:latest AS final
EXPOSE 80 443
LABEL maintainer="z4vr <contact@z4vr.dev>"

WORKDIR /app

COPY --from=web /build/dist/ .
COPY --from=golang /build/shitsuji /bin/shitsuji

RUN chmod +x /bin/shitsuji

ENTRYPOINT [ "/bin/shitsuji" , "-config" , "/etc/shitsuji.toml" ]