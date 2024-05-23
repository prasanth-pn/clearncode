package di

import (
	bootserver "github.com/prasanth-pn/cleancode/pkg/boot"
	"github.com/prasanth-pn/cleancode/pkg/config"
	"github.com/prasanth-pn/cleancode/pkg/db"
	userauth "github.com/prasanth-pn/cleancode/pkg/user_auth"
)

func InitiallizeEvent(conf config.Config) (*bootserver.ServerHttp, error) {

	DB := db.ConnectPGDB(conf)

	userRepository := userauth.NewRepository(DB)
	userService := userauth.NewService(userRepository)
	userHandler := userauth.NewHandler(userService)

	serverHttp := bootserver.NewServerHttp(*userHandler)

	return serverHttp, nil

}
