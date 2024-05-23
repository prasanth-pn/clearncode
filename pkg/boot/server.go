package bootserver

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/cleancode/pkg/config"
	userauth "github.com/prasanth-pn/cleancode/pkg/user_auth"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(userHandler userauth.Handler) *ServerHttp {
	engine := gin.New()
	userHandler.MountRoutes(engine)


	return &ServerHttp{engine}
}

func (s *ServerHttp) Start(conf config.Config) {
	s.engine.Run(conf.Host + ":" + conf.ServerPort)
}
