package pkg

import (
	"apigateway/config"
	user "apigateway/genproto/register"
	stadium "apigateway/genproto/stadium"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStadiumClient(cfg *config.Config) stadium.StadiumServiceClient {
	conn, err := grpc.NewClient(cfg.STADIUM_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return stadium.NewStadiumServiceClient(conn)
}

func NewUserClient(cfg *config.Config) user.RegisterServiceClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return user.NewRegisterServiceClient(conn)
}