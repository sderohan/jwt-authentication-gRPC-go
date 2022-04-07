package client

import (
	"context"
	"time"

	"github.com/sderohan/jwt-authentication-gRPC-go/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{
		service,
		username,
		password,
	}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.AuthServiceClient.Login(ctx, req)
	if err != nil {
		return "", nil
	}

	return res.GetAccessToken(), nil
}
