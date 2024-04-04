## 📬  kvdbstore
HTTP server that can persistently store and retrieve key-value pairs. Exposes a resftull API to interact with stored values and to store new key-value pairs.


## TODO:

1. Create a interface for storage

2. Setup some unit tests 

3. Refactor everything regarding err handling

## Configuration
All configuration is specified in config.yaml

```yaml
server:
  port: "8090"
  host: "localhost"
```

## Running kvdbstore

### Natively using go

entry point resides under cmd/main/main.go, you can chose to run directly towards main.go


```bash
go run cmd/main/main.go
```

or to build a binary
```bash
go build cmd/main/main.go
./main.go
```

or using any containerization tool towards Containerfile, (example with podman)
```bash
podman build .
podman run -it <CONTAINER-ID>
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

### /api/ POST
Store a key-value pair. Asumes a data payload with key and value, see data.json below as example. Returns a 200 OK response and the key-value pair after successfull storeage.

Call:
```bash
curl -X POST -H "Content-Type: application/json" -d @data.json http://localhost:8090/api/some_key
```

data.json:
```json
{ "value" : "some_value" }
```

Response:
```json
{"some_key":"some_value","status":"200 OK"}
```

### /api/ GET
List the value for a given key. Returns a 200 OK response and the value of the corresponding key.

Call:
```bash
curl localhost:8090/api/list/some_key
```

Response:
```json
{"data":"some_value","status":"200 OK"}
```
