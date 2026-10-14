package server

import (
	"context"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/api"
	remoteapi "github.com/Chriss-Leo/qudaju-mprogram/backend/common/api/remote"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine, deps *Dependency) {
	v1 := r.Group("/v1")
	// identity_ctl.go: public exchange; exact registered route only.
	identityGroup := v1.Group("")
	{
		identityGroup.POST(api.WeChatLogin, middleware.LoginRateLimit(deps.IdentitySvc.AllowLogin), middleware.Audit("wechat_login"), controller.WeChatLogin(deps.IdentitySvc))
	}
	auth := middleware.RequireLogin(func(ctx context.Context, token string) (middleware.Principal, error) {
		user, err := deps.IdentitySvc.Me(ctx, token)
		if err != nil {
			return middleware.Principal{}, err
		}
		return middleware.Principal{UserID: user.ID, RoleCodes: []string{"member"}}, nil
	})
	// identity_ctl.go: authentication -> role -> audit -> handler.
	protectedGroup := v1.Group("", auth)
	{
		protectedGroup.GET(api.AuthMe, middleware.ApiPermission("member"), controller.Me(deps.IdentitySvc))
		protectedGroup.POST(api.AuthLogout, middleware.ApiPermission("member"), middleware.Audit("logout"), controller.Logout(deps.IdentitySvc))
	}
	// health_ctl.go: infrastructure probes, no user data.
	healthGroup := r.Group("")
	{
		healthGroup.GET(api.HealthLive, controller.Live())
		healthGroup.GET(api.HealthReady, controller.Ready(deps.Ready))
	}
}
func RegisterRemoteRoutes(r *gin.Engine, deps *Dependency) {
	_ = deps
	remoteRouter := r.Group(remoteapi.Prefix)
	_ = remoteRouter /* no remote APIs; X-Internal-Call never bypasses authentication */
}
