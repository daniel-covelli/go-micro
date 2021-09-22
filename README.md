# learn-go

[![Go Report Card](https://goreportcard.com/badge/github.com/daniel-covelli/learn-go)](https://goreportcard.com/report/github.com/daniel-covelli/learn-go) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a playground for learning microservices in GO.

```shell
# GET
curl http://localhost:9090 | jq # ping server and display formatted json

# POST
curl http://localhost:9090 -d "{...}" | jq

# PUT
curl -v http://localhost:9090/1 -XPUT -d '{...}' | jq
```

## Contents

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

## Tech

Below are the tools and technologies covered in the course.

| Tag                                                             | Description                                                                                                                |
| --------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| [net/http](https://pkg.go.dev/net/http@go1.17.1)                | Package http provides HTTP client and server implementations.                                                              |
| [Gorilla](https://github.com/gorilla/mux)                       | Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler. |
| [Package Validator](https://github.com/go-playground/validator) | Package validator implements value validations for structs and individual fields based on tags.                            |
| [goswagger](https://goswagger.io/)                              | Goswagger autogenerates documentation for Go APIs.                                                                         |
| [Redoc](https://github.com/Redocly/redoc)                       | Redoc is a OpenAPI/Swagger-generated API Reference.                                                                        |
