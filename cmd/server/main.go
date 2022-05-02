package main

import (
	"log"
	"net"

	"github.com/sderohan/jwt-authentication-gRPC-go/config"
	"github.com/sderohan/jwt-authentication-gRPC-go/pb"
	"github.com/sderohan/jwt-authentication-gRPC-go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	config.InitConfig()
	authConfig := config.GetAuthConfig()
	jwtManager := service.NewJWTManager(authConfig.SecretKey, authConfig.RefreshDuration)
	authServer := service.NewAuthServer(jwtManager, userStore)
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	reflection.Register(grpcServer)

	serverConfig := config.GetServerConfig()
	listener, err := net.Listen(serverConfig.Protocol, serverConfig.Address+":"+serverConfig.Port)
	if err != nil {
		log.Fatal("Cannot start tcp server : ", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Cannot start gRPC server : ", err)
	}
}
