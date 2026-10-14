package middleware

import (
	"context"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type Principal struct {
	UserID    string
	RoleCodes []string
}
type Authenticator func(context.Context, string) (Principal, error)

const PrincipalKey = "identity-principal"

func Bearer(c *gin.Context) string {
	parts := strings.Fields(c.GetHeader("Authorization"))
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}
	return parts[1]
}
func RequireLogin(auth Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		principal, err := auth(c.Request.Context(), Bearer(c))
		if err != nil || principal.UserID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorize_failed", "error_description": "请重新登录"})
			return
		}
		c.Set(PrincipalKey, principal)
		c.Next()
	}
}
func ApiPermission(codes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get(PrincipalKey)
		p, ok := value.(Principal)
		if exists && ok {
			for _, actual := range p.RoleCodes {
				for _, wanted := range codes {
					if wanted != "" && actual == wanted {
						c.Next()
						return
					}
				}
			}
		}
		request.Fail(c, errcode.NoAuth)
	}
}
func LoginRateLimit(allow func(context.Context, string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !allow(c.Request.Context(), c.ClientIP()) {
			request.Fail(c, errcode.RateLimited)
			return
		}
		c.Next()
	}
}
func Audit(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		value, _ := c.Get(PrincipalKey)
		p, _ := value.(Principal)
		log.WithFields(log.Fields{"action": action, "actor_id": p.UserID, "route": c.FullPath(), "http_status": c.Writer.Status(), "duration_ms": time.Since(start).Milliseconds()}).Info("audit")
	}
}
func RegisterMdl(r *gin.Engine, origins []string) {
	r.Use(UnKnowErrorHandler())
	r.Use(func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4096)
		c.Header("Cache-Control", "no-store")
		c.Header("X-Content-Type-Options", "nosniff")
		origin := c.GetHeader("Origin")
		if origin != "" {
			allowed := false
			for _, value := range origins {
				if value == origin {
					allowed = true
					break
				}
			}
			if !allowed {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept-Language")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
}
