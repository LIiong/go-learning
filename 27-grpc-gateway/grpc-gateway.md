> grpc-gateway 用于同时生成go 的rpc和HTTP接口
>
> 引用文档：https://github.com/grpc-ecosystem/grpc-gateway

* 安装插件

```go
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

* 编写`helloword.proto`

```protobuf
syntax = "proto3";
package helloworld;
option go_package = "go-learning/27-grpc/proto/helloworld";

// The greeting service definition
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

* 编写grpc-gateway配置`http.yaml`,用于生成HTTP 接口，参考[官方用例](https://github.com/GoogleCloudPlatform/python-docs-samples/blob/master/endpoints/bookstore-grpc/api_config_http.yaml)

```yaml
type: google.api.Service
config_version: 3

http:
  rules:
    - selector: helloworld.Greeter.SayHello
      post: /v1/example/echo
      body: "*"
```

* 编写脚本`gen.sh`,运行脚本生成go文件

```shell
PROTO_PATH=./proto/helloworld
GO_OUT_PATH=./proto/helloworld

protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH helloworld.proto
protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH helloworld.proto
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/http.yaml:$GO_OUT_PATH helloworld.proto
```

* 编写main函数，启动 `go run main.go`
* Curl 请求HTTP接口

```shell
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'
```

