package controller

import (
	"context"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"github.com/gin-gonic/gin"
	"time"
)

func Live() gin.HandlerFunc { return func(c *gin.Context) { request.OK(c, gin.H{"live": true}) } }
func Ready(check func(context.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
		defer cancel()
		if check(ctx) != nil {
			request.OKWithCode(c, 503, gin.H{"ready": false})
			return
		}
		request.OK(c, gin.H{"ready": true})
	}
}
