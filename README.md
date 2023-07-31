## quickstart

```bash
go run main.go
```

## test

```bash
go run client.go
go run test_subscribe.go
go run test_delete.go
```

## update proto

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    clipboard_service/clipboard_service.proto
```