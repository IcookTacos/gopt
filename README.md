## 📬  gottp
A http server written in go that can be used to persistently store key value pairs. Knowing the key of a key-value pair can be used to query the server for it's corresponding value.

## Table of Contents
1. [Configuration](#Configuration)
2. [Endpoints](#Endpoints)
   1. [/api/status](#/api/status-21)
   2. [/api/store](#/api/store-22)

## Configuration
All configuration is specified in config.yaml

```yaml
server:
  port: "8090"
  host: "localhost"
```

The repo comes packaged with a pre-built binary, if needed to run / build  CGO might be needed

```bash
export CGO_ENABLED=1
```

## Endpoints
Configure the host and port you wish to run the server on. Below examplees assumes localhost and 8090.

### /api/status
This endpoint is meant to be used as a basic health check / sanity check. It holds little to no logic and sole purpose is to test connectivity towards the server and that the HTTP call / response flow is ensured.

Call:
```bash
curl localhost:8090/api/status
```

Response:
```json
{"data":"\n","status":"200 OK\n"}
```

### /api/store
Making a POST request towrads this endpoint with a key-value pair and the server will store these.

Call:
```bash
curl -X POST -H "Content-Type: application/json" -d @data.json http://localhost:8090/api/store
```

data.json:
```json
{ "key" : "your_key", "value" : "some_value" }
```

Response:
```json
{"your_key":"some_value","status":"200 OK"}
```
