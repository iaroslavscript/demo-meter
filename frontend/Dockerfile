# syntax = docker/dockerfile:1.2
FROM nginx:1.19.7-alpine

LABEL author="Iaroslav Akimov"

# server http port
EXPOSE 8080

RUN --mount=type=bind,target=/usr/src/app \
    mkdir -p /etc/nginx/templates/ \
    && rm -rf /etc/nginx/conf.d/* /usr/share/nginx/html/* /etc/nginx/templates/* \
    && cp /usr/src/app/nginx/templates/* /etc/nginx/templates/ \
    && cp /usr/src/app/html/* /usr/share/nginx/html/
