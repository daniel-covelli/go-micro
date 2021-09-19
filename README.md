# learn-go

This is a playground for learning microservices in GO.

```shell
# GET
curl http://localhost:9090 | jq # ping server and display formatted json

# POST
curl http://localhost:9090 -d "{...}" | jq

# PUT
curl -v http://localhost:9090/1 -XPUT -d '{...}' | jq
```

## Content

Content in this repo comes from [Nic Jackson](https://github.com/nicholasjackson)'s Building Microservices in Go YouTube [playlist](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_).

- Introduction to microservices
- RESTFul microservices
- gRPC microservices
- Packaging applications with Docker
- Testing microservice
- Continuous Delivery
- Observability
- Using Kubernetes
- Debugging
- Security
- Asynchronous microservices
- Caching
- Microservice reliability using a Service Mesh
