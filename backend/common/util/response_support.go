package util

import (
	"context"
	"encoding/json"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/cache"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// GetLanguage 是独立模板适配：支持 Gin 请求头及项目标准 context 语言值。
// 不解析 Accept-Language 多语言权重；沿用单标签、按连字符取语言的约定。
func GetLanguage(ctx context.Context) string {
	lan, _ := ctx.Value("Accept-Language").(string)
	if c, ok := ctx.(*gin.Context); ok {
		if value, ok := c.Request.Context().Value("Accept-Language").(string); ok {
			lan = value
		}
		if lan == "" {
			lan = c.GetHeader("Accept-Language")
		}
	}
	if lan == "" {
		return "en"
	}
	return strings.Split(lan, "-")[0]
}

// GetI18nMessage 从统一 Hash 取文案。模板为自包含而使用标准 JSON 解码；
// 目标 common 已有该函数时使用原函数，不重复复制此文件。
func GetI18nMessage(ctx context.Context, code string) string {
	if code == "" || cache.RDB == nil {
		return ""
	}
	redisCtx := ctx
	if c, ok := ctx.(*gin.Context); ok {
		redisCtx = c.Request.Context()
	}
	msg, err := cache.RDB.HGet(redisCtx, "i18nMessage", code).Result()
	if err != nil {
		log.WithError(err).Error("GetI18nMessage failed")
		return ""
	}
	var values map[string]interface{}
	if err := json.Unmarshal([]byte(msg), &values); err != nil {
		log.WithError(err).Error("GetI18nMessage decode failed")
		return ""
	}
	// 与来源不同：安全处理缺失或非字符串语言字段，避免类型断言 panic。
	text, _ := values[GetLanguage(ctx)].(string)
	if strings.TrimSpace(text) == "" {
		text, _ = values["en"].(string)
	}
	return text
}

func FormatJavaDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
