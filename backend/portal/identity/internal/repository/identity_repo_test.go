package repository

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/db"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"github.com/spf13/viper"
)

// Opt in against a local development database only. All test rows roll back;
// PostgreSQL identity sequences may advance, as with any rolled-back insert.
func TestPhoneUpsertPostgres(t *testing.T) {
	file := os.Getenv("QUDADJU_LOCAL_DB_TEST_CONFIG")
	if file == "" {
		t.Skip("local PostgreSQL test not requested")
	}
	v := viper.New()
	v.SetConfigFile(file)
	if v.ReadInConfig() != nil {
		t.Fatal("local test configuration unavailable")
	}
	connection, err := db.Open(v.GetString("postgres.dsn"))
	if err != nil {
		t.Fatal("local PostgreSQL unavailable")
	}
	sql, _ := connection.DB()
	defer sql.Close()
	tx := connection.Begin()
	if tx.Error != nil {
		t.Fatal("test transaction failed")
	}
	defer tx.Rollback()
	repo := NewIdentityRepo(tx)
	app, open := "test-phone-registration", fmt.Sprintf("test-%d", time.Now().UnixNano())
	ctx := context.Background()
	first, err := repo.FindOrCreatePhone(ctx, app, open, "+8613800000000")
	if err != nil || first.ID <= 0 || first.PhoneNumber == nil || first.PhoneVerifiedTime == nil {
		t.Fatal("phone registration failed")
	}
	if tx.Model(&model.User{}).Where("id = ?", first.ID).Update("nickname", "KeepName").Error != nil {
		t.Fatal("test setup failed")
	}
	again, err := repo.FindOrCreatePhone(ctx, app, open, "+8613900000000")
	if err != nil || again.ID != first.ID || again.Nickname != "KeepName" || again.PhoneNumber == nil || *again.PhoneNumber != "+8613900000000" {
		t.Fatal("phone login must preserve identity and profile")
	}
	if tx.Model(&model.User{}).Where("id = ?", first.ID).Update("status", 0).Error != nil {
		t.Fatal("test setup failed")
	}
	disabled, err := repo.FindOrCreatePhone(ctx, app, open, "+8613700000000")
	if err != nil || disabled.Status != 0 || *disabled.PhoneNumber != "+8613900000000" {
		t.Fatal("disabled user mutated by login")
	}
	if tx.Model(&model.User{}).Where("id = ?", first.ID).Update("deleted_at", time.Now()).Error != nil {
		t.Fatal("test setup failed")
	}
	if _, err := repo.FindOrCreatePhone(ctx, app, open, "+8613600000000"); err == nil {
		t.Fatal("deleted identity re-registered")
	}
}
