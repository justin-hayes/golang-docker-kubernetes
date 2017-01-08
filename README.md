A "hello world" project created to explore the development environment setup, dependencies, and basic workflow of deploying microservices to Kubernetes.

## Dependencies

- Go
  - Negroni
  - Pat
- Glide
- Docker
- Kubernetes

## Build and deploy

From a fresh clone of the repository:
```
glide install
docker build -t golang-docker-kubernetes .
docker run -p 8080:8080 --rm --name gdk golang-docker-kubernetes
```
The application will be accessible on port 8080.

Todo:
- Kubernetes
