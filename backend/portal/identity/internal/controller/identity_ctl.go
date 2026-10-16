package controller

import (
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model/dto"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/service"
	"github.com/gin-gonic/gin"
)

func WeChatLogin(svc *service.IdentitySvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.WeChatLogin
		if err := c.ShouldBindJSON(&input); err != nil {
			request.Fail(c, errcode.ParameterError)
			return
		}
		result, err := svc.Login(c.Request.Context(), input.Code, input.PhoneCode)
		if err != nil {
			request.Fail(c, err)
			return
		}
		c.Set(middleware.PrincipalKey, middleware.Principal{UserID: result.User.ID, RoleCodes: []string{"member"}})
		request.OK(c, result)
	}
}
func Me(svc *service.IdentitySvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := svc.Me(c.Request.Context(), middleware.Bearer(c))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "authorize_failed", "error_description": "请重新登录"})
			return
		}
		request.OK(c, user)
	}
}
func Logout(svc *service.IdentitySvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := svc.Logout(c.Request.Context(), middleware.Bearer(c)); err != nil {
			request.Fail(c, err)
			return
		}
		request.OKEmpty(c)
	}
}
