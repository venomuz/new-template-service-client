package main

import (
	"context"
	"fmt"
	api "github.com/venomuz/new-template-service-api/api"
	"github.com/venomuz/new-template-service-client/configs"
	"github.com/venomuz/new-template-service-client/pkg/logger"
	"google.golang.org/grpc"
)

func main() {

	cfg := configs.Load()

	log := logger.New(cfg.LogLevel, "template-go-service-client")

	grpcTarget := fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(grpcTarget, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connect to grpc server error", logger.Error(err))
	}
	defer conn.Close()

	client := api.NewUsersServiceClient(conn)

	create, err := client.Create(context.Background(), &api.UserRequest{
		RegionId:    1,
		Role:        "asdasd",
		Username:    "test",
		Password:    "test",
		FullName:    "test",
		Phone:       "test",
		Email:       "test",
		Description: "test",
		State:       1,
	})

	if err != nil {
		log.Error("error while create", logger.Error(err))
	}
	fmt.Println(create)
}
