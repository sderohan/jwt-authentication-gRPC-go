package main

import (
	"log"
	"net"
	"time"

	"github.com/sderohan/jwt-authentication-gRPC-go/pb"
	"github.com/sderohan/jwt-authentication-gRPC-go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// SECRET should be stored securely
	SECRET         = "strong-secret-message"
	TOKEN_DURATION = 15 * time.Minute
)

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func accessibleRoles() map[string][]string {
	// populate the below map as per the user roles
	return map[string][]string{}
}

func main() {
	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("Cannot seed users ", err)
	}

	jwtManager := service.NewJWTManager(SECRET, TOKEN_DURATION)
	authServer := service.NewAuthServer(jwtManager)
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:56878")
	if err != nil {
		log.Fatal("Cannot start tcp server : ", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Cannot start gRPC server : ", err)
	}
}
