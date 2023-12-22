## 📬  gottp
A http server written in go that can be used to persistently store key value pairs. Knowing the key of a key-value pair can be used to query the server for it's corresponding value.

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

## API & Endpoints
Configure the host and port you wish to run the server on. Below examplees assumes localhost and 8090.

### /api/status

Example call:
```bash
curl localhost:8090/api/status
```

Response:
```json
{"data":"\n","status":"200 OK\n"}
```

### /api/store

Example call:
```bash
curl -X POST -H "Content-Type: application/json" -d @data.json http://localhost:8090/api/store
```

Expected data.json:
```json
{
    "key" : "your_key",
    "value" : "some_value"
}
```

Response:
```json
{"data":"\n","status":"200 OK\n"}
{"your_key":"some_value","status":"200 OK"}
```
