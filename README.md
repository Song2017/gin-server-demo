# Gin Server
Gin Server

### Features
1. Swagger UI
2. Docker
3. Prometheus
4. Shell generator

### Production
1. Cypher Service
And we can get the authorization key SECURITY_CA_KEY in Vault.
```shell script
curl -X 'POST' \
  'http://localhost:8000/platform/cypher?Authorization=SECURITY_CA_KEY' \
  -H 'Content-Type: application/json' \
  -d '{
  "storeId": "platform.test",
  "operation": "encrypt_batch",
  "items": [
    "test"
  ]
}'
```

## Usage Locally
Go package
```
go mod init api_server/v1
export GOPROXY=https://goproxy.io
go mod tidy
```

To run the server, please execute the following from the root directory:
```
source bin/gen_go_server.sh
source bin/run_swagger_server.sh
```

and open your browser to here:
```
http://localhost:8000/swagger/index.html
```

Your OpenAPI definition lives here:
```
http://localhost:8000/openapi.yaml
```

Docker image
Please add IP to PYCredit white list or enable test VPN
```
docker build -t nomad:test .
docker run -d --cap-add NET_ADMIN --device /dev/net/tun -p 8000:8000 nomad:test

curl --location --request GET \
'http://localhost:8000/cbec/id?user_id=123**&name=宋**&Authorization=test'
```
