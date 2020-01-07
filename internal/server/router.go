package server

import (
	"github.com/gin-gonic/gin"
	"wdkj/account/internal/service"
)

type Server struct {
	loginManger service.LoginManager
}

func(s *Server)newRouter () {
	r := gin.Default()
	g := r.Group("/api/v1")
	g.POST("login", s.Login)
}