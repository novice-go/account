package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	serverModel "wdkj/account/internal/server/server-model"
	"wdkj/account/model"
	"wdkj/account/utils"
	"wdkj/account/utils/log"
)

// 登录
//func (s *Server) Login(c *gin.Context) {
//	// TODO get param
//	if err := s.loginManger.Login(); err != nil {
//		// TODO log
//		utils.HttpResponseErr(c, err)
//		return
//	}
//
//	utils.HttpResponseSuccess(c, nil)
//}
//
// 注册
//func (s *Server) Register(c *gin.Context) {
//	// TODO get apram
//	if err := s.loginManger.Register(); err != nil {
//		// TODO log
//		utils.HttpResponseErr(c, err)
//		return
//	}
//
//	utils.HttpResponseSuccess(c, nil)
//}

// 生成验证码
func (s *Server) GenVCode(c *gin.Context) {
	var req serverModel.GenVCodeReq

	if err := c.BindJSON(&req); err != nil {
		log.Logger.Error("bind json fail", zap.Error(err))
		utils.HttpResponseErr(c, err)
		return
	}

	if err := req.IsValid(); err != nil {
		log.Logger.Error("req valid", zap.Error(err))
		utils.HttpResponseErr(c, err)
		return
	}

	ctx := &model.VCodeContext{
		VCodeType: req.VCodeType,
		Phone:     req.Phone,
	}
	if err := s.loginManger.GenVCode(ctx); err != nil {
		log.Logger.Error("gen vcode fail", zap.Any("ctx", ctx), zap.Error(err))
		utils.HttpResponseErr(c, err)
		return
	}

	utils.HttpResponseSuccess(c, nil)
}
