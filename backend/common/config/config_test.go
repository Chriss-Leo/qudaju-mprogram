package config

import (
	"strings"
	"testing"
	"time"
)

func TestSecretsRequired(t *testing.T) {
	var c Config
	c.Server.Port = 8081
	c.Auth.TTL = time.Hour
	c.Postgres.DSN = "test-dsn"
	c.Postgres.Schema = "qudaju_identity"
	c.Redis.Addr = "localhost:6379"
	c.Log.Level = "info"
	err := c.Validate()
	if err == nil || strings.Contains(err.Error(), "test-dsn") {
		t.Fatal("secret validation or redaction failed")
	}
	c.WeChat.AppID = "test-app"
	c.WeChat.AppSecret = "test-secret"
	if err = c.Validate(); err != nil {
		t.Fatal(err)
	}
	c.Auth.TTL = 0
	if c.Validate() == nil {
		t.Fatal("zero TTL accepted")
	}
}
