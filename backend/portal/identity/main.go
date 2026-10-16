package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/cache"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/config"
	dbcommon "github.com/Chriss-Leo/qudaju-mprogram/backend/common/db"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/starter"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/server"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.WithField("reason", err.Error()).Error("identity startup stopped")
		os.Exit(1)
	}
}
func run() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}
	config.SetLogger(conf)
	gin.SetMode(gin.ReleaseMode)
	db, err := dbcommon.Open(conf.Postgres.DSN)
	if err != nil {
		return err
	}
	sql, err := db.DB()
	if err != nil {
		return errors.New("database handle unavailable")
	}
	defer sql.Close()
	if !db.Migrator().HasTable(&model.User{}) {
		return errors.New("identity migration missing")
	}
	client := redis.NewClient(&redis.Options{Addr: conf.Redis.Addr, Password: conf.Redis.Password, DB: conf.Redis.DB, DialTimeout: 5 * time.Second, ReadTimeout: 3 * time.Second, WriteTimeout: 3 * time.Second})
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = client.Ping(ctx).Err(); err != nil {
		return errors.New("redis unavailable")
	}
	type language struct {
		Code string
		EN   string `gorm:"column:en"`
		ZH   string `gorm:"column:zh"`
	}
	var rows []language
	if err = db.WithContext(ctx).Table("public.base_language").Find(&rows).Error; err != nil || len(rows) == 0 {
		return errors.New("language resources missing; apply migrations first")
	}
	pipeline := client.TxPipeline()
	for _, row := range rows {
		if row.EN == "" {
			return errors.New("language fallback missing")
		}
		message, err := json.Marshal(map[string]string{"en": row.EN, "zh": row.ZH})
		if err != nil {
			return errors.New("language resource encoding failed")
		}
		pipeline.HSet(ctx, "i18nMessage", row.Code, string(message))
	}
	if _, err = pipeline.Exec(ctx); err != nil {
		return errors.New("language warmup failed")
	}
	cache.RDB = client
	root, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	srv := server.New(conf, db, client)
	log.WithField("service", conf.Server.Name).Info("dependencies ready; starting listener")
	return starter.EnableHTTP(root, fmt.Sprintf(":%d", conf.Server.Port), srv)
}
