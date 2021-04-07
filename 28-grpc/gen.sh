PROTO_PATH=./helloworld
GO_OUT_PATH=./helloworld

protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH helloworld.proto
protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH helloworld.proto