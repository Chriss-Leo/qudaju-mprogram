package server

import (
	"context"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/config"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/rpc"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/starter"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/repository"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type IdentityServer struct{ GSrv *gin.Engine }
type Dependency struct {
	IdentitySvc *service.IdentitySvc
	Ready       func(context.Context) error
}

func New(conf *config.Config, db *gorm.DB, redis redis.UniversalClient) starter.Server {
	svc := service.NewIdentitySvc(repository.NewIdentityRepo(db), repository.NewSessionRepo(redis), rpc.NewWeChatClient(conf.WeChat.AppID, conf.WeChat.AppSecret), conf.WeChat.AppID, conf.Auth.TTL)
	deps := Dependency{IdentitySvc: svc, Ready: func(ctx context.Context) error {
		sql, err := db.DB()
		if err != nil {
			return err
		}
		if err = sql.PingContext(ctx); err != nil {
			return err
		}
		return redis.Ping(ctx).Err()
	}}
	r := gin.New()
	r.ContextWithFallback = true
	_ = r.SetTrustedProxies(nil)
	middleware.RegisterMdl(r, conf.Server.Origins)
	RegisterApiRoutes(r, &deps)
	RegisterRemoteRoutes(r, &deps)
	return &IdentityServer{GSrv: r}
}
func (s *IdentityServer) GetGinSrv() *gin.Engine { return s.GSrv }
