# learn-go

[![Go Report Card](https://goreportcard.com/badge/github.com/daniel-covelli/learn-go)](https://goreportcard.com/report/github.com/daniel-covelli/learn-go)

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

The contents of this repo come from [Nic Jackson](https://github.com/nicholasjackson)'s [Building Microservices in Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) YouTube series. The concepts below are covered in the full 21 part series.

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
