
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
