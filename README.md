## 📬  kvdbstore
A http server written in go that can be used to persistently store & retrieve key value pairs.

## Contents:
1. [Configuration](#Configuration)
2. [Endpoints](#Endpoints)
   1. [/api/status](#apistatus)
   2. [/api/store](#apistore)

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
Ensures basic functionality and can be used to test connectivity. Returns a 200 OK response on request.

Call:
```bash
curl localhost:8090/api/status
```

Response:
```json
{"data":"\n","status":"200 OK\n"}
```

### /api/store
Store a key-value pair. Asumes a data payload with key and value, see data.json below as example. Returns a 200 OK response and the key-value pair after successfull storeage.

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

### /api/list
List the value for a given key. Returns a 200 OK response and the value of the corresponding key.

Call:
```bash
curl localhost:8090/api/list/your_key
```

Response:
```json
{"data":"some_value","status":"200 OK"}
```
