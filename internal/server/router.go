package server

import (
	"github.com/gin-gonic/gin"
	"wdkj/account/internal/service"
)

type Server struct {
	loginManger *service.LoginManager
}

func (s *Server) NewRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	g := e.Group("/api/v1/")
	//g.POST("login", s.Login)
	//g.POST("register", s.Register)
	g.POST("gen_vcode", s.GenVCode)

	return e
}
