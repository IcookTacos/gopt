## 📬  gottp
A http server written in go

## 📂 Configuration & Other
All configuration is specified in config.yaml

The repo comes packaged with a pre-built binary, if needed to run / build  CGO might be needed
```bash
export CGO_ENABLED=1
```

## 📧 Usage
Configure the host and port you wish to run the server on.

```bash
go run .
```

Example using curl as a client to interact with the server 
```bash
curl localhost:8090/api/status
```

Example response
```json
{"data":"\n","status":"200 OK\n"}
```
