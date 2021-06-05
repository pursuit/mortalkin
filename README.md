# mortalkin

## Development Guide
### Pre-requisite
- [Go 1.16](https://golang.org/doc/install)

### Generate Proto
```
protoc --proto_path=pkg/proto --go_out=internal/proto/out --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=internal/proto/out user.proto
```
