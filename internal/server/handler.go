package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func (s *Server) Login(c *gin.Context) {
	if err := s.loginManger.Login(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}