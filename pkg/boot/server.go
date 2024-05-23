package bootserver

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/cleancode/pkg/config"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp() *ServerHttp {
	engine := gin.New()
	return &ServerHttp{engine}
}

func (s *ServerHttp) Start(conf config.Config) {
	s.engine.Run(conf.Host + ":" + conf.ServerPort)
}
