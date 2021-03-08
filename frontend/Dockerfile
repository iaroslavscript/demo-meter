# syntax = docker/dockerfile:1.2

ARG flovar=buster

FROM debian:$flovar

LABEL author="Iaroslav Akimov"

RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt/lists \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    && apt-get update \
    && apt-get install -y --no-install-recommends nginx-full libnginx-mod-http-lua gettext-base

# server http port
EXPOSE 8080

RUN --mount=type=bind,target=/usr/src/app \
    mkdir -p /etc/nginx/conf.d \
    && mkdir -p /etc/nginx/templates/ \
    && mkdir /docker-entrypoint.d \
    && cp /usr/src/app/docker-entrypoint.sh / \
    && cp /usr/src/app/docker-entrypoint.d/* /docker-entrypoint.d/ \
    && chmod 0755 /docker-entrypoint.sh \
    && chmod 0755 /docker-entrypoint.d/* \
    && rm -rf /etc/nginx/conf.d/* /usr/share/nginx/html/* /etc/nginx/templates/* \
    && cp /usr/src/app/nginx/nginx.conf /etc/nginx/ \
    && cp /usr/src/app/nginx/templates/* /etc/nginx/templates/ \
    && cp /usr/src/app/html/* /usr/share/nginx/html/

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]