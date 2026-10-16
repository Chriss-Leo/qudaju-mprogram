package config

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	Server struct {
		Name    string   `mapstructure:"name"`
		Port    int      `mapstructure:"http_port"`
		Origins []string `mapstructure:"origins"`
	} `mapstructure:"server"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
	Postgres struct {
		DSN    string `mapstructure:"dsn"`
		Schema string `mapstructure:"schema"`
	} `mapstructure:"postgres"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`
	WeChat struct {
		AppID     string `mapstructure:"app_id"`
		AppSecret string `mapstructure:"app_secret"`
	} `mapstructure:"wechat"`
	Auth struct {
		TTL time.Duration `mapstructure:"session_ttl"`
	} `mapstructure:"auth"`
	Media struct {
		AvatarDir string `mapstructure:"avatar_dir"`
	} `mapstructure:"media"`
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env != "" && !regexp.MustCompile(`^[a-z0-9_-]+$`).MatchString(env) {
		return nil, errors.New("invalid APP_ENV")
	}
	file := "./config/config.yaml"
	if env != "" {
		file = "./config/config." + env + ".yaml"
	}
	v := viper.New()
	v.SetConfigFile(file)
	v.SetDefault("server.name", "qudaju-identity")
	v.SetDefault("server.http_port", 8081)
	v.SetDefault("log.level", "info")
	v.SetDefault("postgres.schema", "public")
	v.SetDefault("auth.session_ttl", "2h")
	v.SetDefault("media.avatar_dir", "./storage/avatars")
	for key, name := range map[string]string{"server.http_port": "APP_HTTP_PORT", "postgres.dsn": "APP_POSTGRES_DSN", "redis.addr": "APP_REDIS_ADDR", "redis.password": "APP_REDIS_PASSWORD", "redis.db": "APP_REDIS_DB", "wechat.app_id": "APP_WECHAT_APP_ID", "wechat.app_secret": "APP_WECHAT_APP_SECRET", "auth.session_ttl": "APP_SESSION_TTL", "log.level": "APP_LOG_LEVEL"} {
		if err := v.BindEnv(key, name); err != nil {
			return nil, errors.New("configuration binding failed")
		}
	}
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.New("configuration file unavailable")
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, errors.New("configuration invalid")
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &c, nil
}
func (c *Config) Validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 || c.Auth.TTL < time.Minute || c.Auth.TTL > 24*time.Hour {
		return errors.New("server port or session TTL invalid")
	}
	if strings.TrimSpace(c.Postgres.DSN) == "" || c.Postgres.Schema != "public" || c.Redis.Addr == "" || c.Redis.DB < 0 || c.Redis.DB > 15 {
		return errors.New("database/redis configuration missing or invalid")
	}
	if c.WeChat.AppID == "" || c.WeChat.AppSecret == "" {
		return errors.New("APP_WECHAT_APP_ID and APP_WECHAT_APP_SECRET are required")
	}
	if _, err := log.ParseLevel(c.Log.Level); err != nil {
		return errors.New("log level invalid")
	}
	return nil
}
func SetLogger(c *Config) {
	level, _ := log.ParseLevel(c.Log.Level)
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
