package main

import (
	"context"
	"fmt"
	"grpc_hello_world/config"
	"grpc_hello_world/controller"
	"grpc_hello_world/infrastructure/dao"
	"grpc_hello_world/model/user"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func IsValidAppKey(appKey []string) bool {
	if len(appKey) < 1 {
		return false
	}

	return appKey[0] == "12345678"
}

func UnaryInterceptorImpl(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("incoming request %s\n", info.FullMethod)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
	}

	if !IsValidAppKey(md["app_key"]) {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated")
	}

	m, err := handler(ctx, req)

	return m, err
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	//Create new gRPC server handler
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptorImpl))

	repo := dao.NewUserRepository(db)

	//register gRPC UserService to gRPC server handler
	user.RegisterUserServiceServer(server, controller.NewUserService(repo))

	reflection.Register(server)

	//Run server
	if err := server.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
