## quickstart

```bash
bazel run //server:server
```

## test

```bash
bazel run //test:test_delete
bazel run //test:test_getclipboards
bazel run //test:test_send
bazel run //test:test_subscribe
```

## for development

### build

```bash
bazel build //server:server
```
cross-compiling
```bash
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 server
````
### generate proto files

**protoc** needs to be installed

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/clipboard_service.proto
```

### update deps

```bash
bazel run //:gazelle
bazel run //:gazelle-update-repos 
```