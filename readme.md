# Small starter for go-grpc-k8s. 

## My Notes on setting up

##### Gets for gRPC and Protobuf
```shell
$ go get google.golang.org/grpc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
$ go get google.golang.org/protobuf/runtime/protoimpl@v1.25.0
$ go get google.golang.org/protobuf/runtime/protoimpl@v1.25.0
```

##### Setup gRPC Proto file
```shell
#Generate the pb files:
protoc --go_out=. --go-grpc_out=. proto/services.proto
```

#### Setup K8s
```shell
# Docker build commands
# Server
docker build . -t williamnoble/add-service:v1.0
docker run -p :3001:3000 williamnoble/add-service:v1.0

# Client
docker push williamnoble/add-service:v1.0
docker build . -t williamnoble/api-service:v1.0
```


Run server/client from their respective directories (vs the root dirs)