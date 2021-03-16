# syntax = docker/dockerfile:1.2

ARG flovar=buster

#FROM debian:$flovar as tools

#RUN --mount=type=bind,target=/usr/src/app \
#    && cp -r /usr/src/app/tools/docker-entrypoint.d / \
#    && cp /usr/src/app/tools/docker-entrypoint.sh / \
#    && cp /usr/src/app/tools/wait-for.sh /usr/local/bin/wait-for \
#    && chmod 0777 /docker-entrypoint.sh \
#    && chmod -R 0777 /usr/local/bin/* \
#    && rm -rf /etc/nginx/conf.d/* /usr/share/nginx/html/* /etc/nginx/templates/* \
#    && cp /usr/src/app/nginx/nginx.conf /etc/nginx/ \
#    && cp /usr/src/app/nginx/templates/* /etc/nginx/templates/ \
#    && cp /usr/src/app/html/* /usr/share/nginx/html/


FROM debian:$flovar

LABEL author="Iaroslav Akimov"

# docker-entrypoint.d depends on gettext-base
# wait-for depends on netcat
RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt/lists \
    rm -f /etc/apt/apt.conf.d/docker-clean \
    && apt-get update \
    && apt-get install -y --no-install-recommends nginx-full libnginx-mod-http-lua gettext-base netcat

# server http port
EXPOSE 8080

RUN --mount=type=bind,target=/usr/src/app \
    cp -r /usr/src/app/tools/docker-entrypoint.d / \
    && cp /usr/src/app/tools/docker-entrypoint.sh / \
    && cp /usr/src/app/tools/wait-for.sh /usr/local/bin/wait-for \
    && chmod 0777 /docker-entrypoint.sh \
    && chmod -R 0777 /docker-entrypoint.d/ \
    && chmod -R 0777 /usr/local/bin/* \
    && mkdir -p /etc/nginx/conf.d \
    && mkdir -p /etc/nginx/templates/ \
    && rm -rf /etc/nginx/conf.d/* /usr/share/nginx/html/* /etc/nginx/templates/* \
    && cp /usr/src/app/nginx/nginx.conf /etc/nginx/ \
    && cp /usr/src/app/nginx/templates/* /etc/nginx/templates/ \
    && cp /usr/src/app/html/* /usr/share/nginx/html/


# this env variables are also used by nginx at runtime, 
# see /etc/nginx/templates/*.template
ENV \
    NGINX_DEFAULT_SERVER_ADDR="0.0.0.0" \
    NGINX_DEFAULT_SERVER_PORT="8080" \
    NGINX_UPSTREAM_BACKEND_SERVER="backend:8080" \
    NGINX_UPSTREAM_BACKEND_SCHEME="http"

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]
