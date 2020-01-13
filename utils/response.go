package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpResp struct {
	Success bool        `json:"success"`
	ErrMsg  string      `json:"err_msg ,omitempty"` // 错误时有值
	Data    interface{} `json:"data,omitempty"`     // 成功时有值
}

func HttpResponseErr(c *gin.Context, e error) {
	errCode := ParseErr(e)
	if errCode != "" {
		c.JSON(http.StatusBadRequest, &HttpResp{
			Success: false,
			ErrMsg:  e.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &HttpResp{
		Success: false,
		ErrMsg:  SysErr.Error(),
	})
}

func HttpResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &HttpResp{
		Success: true,
		Data:    data,
	})
}
