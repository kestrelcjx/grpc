package main

import (
	"log"
	"net"

	"github.com/kestrel/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

// protoc
// proto-gen-go
// proto-gen-go-grpc
// option go_package = ".;xx"
// mkdir xx
// protoc -proto_path=proto xx.proto -go_out=xx -go-grpc_out=xx
// grpcurl -plaintext -d '{\"key\": \"value\"}' ip:port xx.xxService.xxMethod
