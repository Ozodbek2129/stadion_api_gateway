package handler

import (
	"apigateway/config"
	user "apigateway/genproto/register"
	stadium "apigateway/genproto/stadium"
	pkg "apigateway/pkg/client"
	"apigateway/pkg/logger"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	UserService user.RegisterServiceClient
	Log         *slog.Logger
	Stadium     stadium.StadiumServiceClient
	Enforcer    *casbin.Enforcer
}

func NewHandler() *Handler {
	cfg := config.Load()
	return &Handler{
		UserService: pkg.NewUserClient(&cfg),
		Log:         logger.NewLogger(),
		Stadium:     pkg.NewStadiumClient(&cfg),
	}
}
