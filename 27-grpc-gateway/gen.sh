PROTO_PATH=./proto/helloworld
GO_OUT_PATH=./proto/helloworld

protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH helloworld.proto
protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH helloworld.proto
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/http.yaml:$GO_OUT_PATH helloworld.proto