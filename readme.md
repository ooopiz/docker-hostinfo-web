# docker-hostinfo-web

The project is a web service running in docker, I use it to test high availability or loading balance architecture.🔨🔨🔨

When you send http request to the container, it will show container host information.

## Usage

Expose port : 9090

### Basic

`$ docker run -itd -p 80:9090 loopbai/hostinfo-in-web`

### Docker swarm

You should add a file (e.g. stack.yml)

```
version: '3.2'
services:
  web:
    image: loopbai/hostinfo-in-web
    ports:
      - 80:9090/tcp
    deploy:
      replicas: 3
```

`$ docker stack deploy -c stack.yml hostinfo`

