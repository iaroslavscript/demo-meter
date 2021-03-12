# syntax = docker/dockerfile:1.2

ARG flovar=alpine

FROM nginx:1.19.7-$flovar

LABEL author="Iaroslav Akimov"

# this env variables are also used by nginx at runtime, 
# see /etc/nginx/templates/*.template
ENV NGINX_DEFAULT_SERVER_ADDR="0.0.0.0"
ENV NGINX_DEFAULT_SERVER_PORT="80"

RUN --mount=type=bind,target=/usr/src/app \
    mkdir -p /etc/nginx/templates/ \
    && rm -rf /etc/nginx/conf.d/* /usr/share/nginx/html/* /etc/nginx/templates/* \
    && cp /usr/src/app/nginx/templates/* /etc/nginx/templates/ \
    && cp /usr/src/app/html/* /usr/share/nginx/html/

VOLUME /etc/nginx/
VOLUME /usr/share/nginx/html/
