
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

### Configuration

Backend accepts enviroment variables listed below:

 * `APP_API_ENTRYPOINT` - address of external API server entrypoint (default: `http://localhost:8000`)
 * `APP_SERVER_BIND_ADDR` - server binding address for HTTP traffic (default: `127.0.0.1`)
 * `APP_SERVER_BIND_PORT` - server binding address for HTTP traffic (default: `8080`)
 * `APP_METRICS_BIND_ADDR` - server binding address for metrics endpoint (default: `127.0.0.1`)
 * `APP_METRICS_BIND_PORT` - server binding port for metrics endpoint (default: `2112`)
