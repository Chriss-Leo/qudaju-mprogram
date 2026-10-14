package middleware

import (
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UnKnowErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 1. 记录未知panic日志
				log.Errorf("panic recovered: %v\n%s", r, string(debug.Stack()))
				// 2. 返回标准 JSON
				request.Fail(c, errcode.Failed)
			}
		}()

		// 继续处理请求
		c.Next()
	}
}
