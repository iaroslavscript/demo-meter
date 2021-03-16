# syntax = docker/dockerfile:1.2

ARG alpine_flovar=alpine3.13

############################
# Alping building stage
############################

FROM golang:1.16.0-$alpine_flovar as builder

RUN --mount=type=cache,target=/etc/apk/cache,mode=0777,rw \
    apk add --update alpine-sdk

RUN --mount=type=cache,target=/root/.cache/go-build,mode=0777,rw \
    --mount=type=bind,target=/usr/src/app \
    cd /usr/src/app && make go-deps

RUN --mount=type=bind,target=/usr/src/app \
    cd /usr/src/app && OUTPUT=/opt/app make build

############################
# Alping release stage
############################

FROM golang:1.16.0-$alpine_flovar AS final
LABEL author="Iaroslav Akimov"

# wait-for depends on netcat
RUN --mount=type=cache,target=/etc/apk/cache,mode=0777,rw\
    apk add netcat-openbsd

RUN --mount=type=bind,target=/usr/src/app \
    cp /usr/src/app/tools/wait-for.sh /usr/local/bin/wait-for \
    && chmod 0777 /usr/local/bin/wait-for

ENTRYPOINT ["/opt/app"]

# server http port
EXPOSE 8080

# server metrics port
EXPOSE 2112

COPY --from=builder /opt/app /opt/app
