
# Frontend

### Building Docker images

Requirements:
* docker engine 19.03

Use `make` for building docker images
```
root@localhost:~# cd frontend && make docker-build
```

For manual build use:
```
root@localhost:~# cd frontend && DOCKER_BUILDKIT=1 docker build --build-arg flovar=alpine -t frontend:latest .
```

# Backend

### Building Docker images

Requirements:
* docker engine 19.03

Use `make` for building docker images
```
root@localhost:~# cd backend && make docker-build
```

For manual build use:
```
root@localhost:~# cd backend && DOCKER_BUILDKIT=1 docker build --build-arg alpine_flovar=alpine3.13 -t backend:latest .
```

# Testing

### Staging API Server

For testing and development purpose there is a simple implementation of API server. It serves static json files stored at `VOLUME /usr/share/nginx/html/`

### Building Docker images

Use `make` for building docker images
```
root@localhost:~# cd apiserver && make docker-build
```

For manual build use:
```
root@localhost:~# cd apiserver && DOCKER_BUILDKIT=1 docker build --build-arg alpine_flovar=alpine -t apiserver:latest .
```

### Configuration

By default API server listens at `localhost:80` but could be configured via enviroment variables at run-time.

Available enviroment variables are listed below:
 * `NGINX_DEFAULT_SERVER_ADDR` default *0.0.0.0*
 * `NGINX_DEFAULT_SERVER_PORT` default *80*

Example:
```
root@localhost:~# docker run -itd -P -e NGINX_DEFAULT_SERVER_PORT='443' apiserver
```
