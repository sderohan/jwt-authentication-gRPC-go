package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sderohan/jwt-authentication-gRPC-go/client"
	"github.com/sderohan/jwt-authentication-gRPC-go/config"
	"google.golang.org/grpc"
)

const (
	username        = "admin1"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {
	// populate the below map to authenticate the RPC's
	return map[string]bool{}
}

func main() {

	transportOption := grpc.WithInsecure()

	// connection specifically maintained for the auth
	cc1, err := grpc.Dial("0.0.0.0:56878", transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	// Register the interceptor on second connection
	// interceptor will utilise the first connection
	cc2, err := grpc.Dial(
		"0.0.0.0:56878",
		transportOption,
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	_ = cc2
	config.InitConfig()
	fmt.Println(config.GetAuthConfig())
	fmt.Println(config.GetServerConfig())
}
