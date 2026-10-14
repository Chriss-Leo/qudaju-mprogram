package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement" json:"-"`
	WeChatAppID  string         `gorm:"column:wechat_app_id" json:"-"`
	WeChatOpenID string         `gorm:"column:wechat_open_id" json:"-"`
	Nickname     string         `gorm:"column:nickname" json:"nickname"`
	AvatarURL    string         `gorm:"column:avatar_url" json:"avatarUrl"`
	Status       int16          `gorm:"column:status" json:"-"`
	CreateUserID int64          `gorm:"column:create_user_id" json:"-"`
	CreateTime   time.Time      `gorm:"column:create_time" json:"-"`
	UpdateUserID int64          `gorm:"column:update_user_id" json:"-"`
	UpdateTime   time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

func (User) TableName() string { return "qudaju_identity.auth_user" }
