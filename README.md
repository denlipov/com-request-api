# Ozon Marketplace Request API

---

## Build project

### Local

For local assembly you need to perform

```zsh
$ make deps # Installation of dependencies
$ make build # Build project
```
## Running

### For local development

```zsh
$ docker-compose up -d
```

---

## Services

### Swagger UI

The Swagger UI is an open source project to visually render documentation for an API defined with the OpenAPI (Swagger) Specification

- http://localhost:8081

### Grafana:

- http://localhost:3000
- - login `admin`
- - password `MYPASSWORT`

### gRPC:

- http://localhost:8082

```sh
grpc_cli call localhost:8082 DescribeRequestV1 'request_id: 101'
grpc_cli call localhost:8082 ListRequestV1 ''
grpc_cli call localhost:8082 CreateRequestV1 'request: { service: "dum", user: "test", text: "hooo" }'
grpc_cli call localhost:8082 RemoveRequestV1 'request_id: 101'
grpc_cli call localhost:8082 UpdateRequestV1 'request_id: 31, body: { text: "some new text", user: "phycus" }'
```

### Gateway:

It reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC

- http://localhost:8080

```sh
curl -s -X GET http://localhost:8080/v1/requests/100 | jq
{
  "value": {
    "id": "100",
    "service": "dummyService",
    "user": "dummyUser",
    "text": "dummyText"
  }
}

curl -s -X POST -d '{"request": {"service": "se", "user": "us", "text": "tx"}}' http://localhost:8080/v1/requests/create | jq
{
  "requestId": "1443635317331776148"
}

curl -s -X GET http://localhost:8080/v1/requests/list | jq
{
  "request": [
    {
      "id": "4751997750760398084",
      "service": "someService1",
      "user": "someUser1",
      "text": "someText1"
    },
    {
      "id": "7504504064263669287",
      "service": "someService2",
      "user": "someUser2",
      "text": "someText2"
    }
  ]
}

curl -s -X DELETE http://localhost:8080/v1/requests/remove/1281233 | jq
{
  "status": true
}
```

### Metrics:

Metrics GRPC Server

- http://localhost:9100/metrics

### Status:

Service condition and its information

- http://localhost:8000
- - `/live`- Layed whether the server is running
- - `/ready` - Is it ready to accept requests
- - `/version` - Version and assembly information

### Prometheus:

Prometheus is an open-source systems monitoring and alerting toolkit

- http://localhost:9090

### Kafka

Apache Kafka is an open-source distributed event streaming platform used by thousands of companies for high-performance data pipelines, streaming analytics, data integration, and mission-critical applications.

- http://localhost:9094

### Kafka UI

UI for Apache Kafka is a simple tool that makes your data flows observable, helps find and troubleshoot issues faster and deliver optimal performance. Its lightweight dashboard makes it easy to track key metrics of your Kafka clusters - Brokers, Topics, Partitions, Production, and Consumption.

- http://localhost:9001

### Jaeger UI

Monitor and troubleshoot transactions in complex distributed systems.

- http://localhost:16686

### Graylog

Graylog is a leading centralized log management solution for capturing, storing, and enabling real-time analysis of terabytes of machine data.

- http://localhost:9000
- - login `admin`
- - password `admin`

### PostgreSQL

For the convenience of working with the database, you can use the [pgcli](https://github.com/dbcli/pgcli) utility. Migrations are rolled out when the service starts. migrations are located in the **./migrations** directory and are created using the [goose](https://github.com/pressly/goose) tool.

```sh
$ pgcli "postgresql://docker:docker@localhost:5432/com_request_api"
```

### Python client

```shell
$ python -m venv .venv
$ . .venv/bin/activate
$ make deps
$ make generate
$ cd pypkg/com-request-api
$ python setup.py install
$ cd ../..
$ docker-compose up -d
$ python scripts/grpc_client.py
```


### Thanks

- [Evald Smalyakov](https://github.com/evald24)
- [Michael Morgoev](https://github.com/zerospiel)
