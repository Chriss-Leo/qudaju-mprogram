package repository

import (
	"context"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/db/repo"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type IdentityRepo struct{ repo.BaseRepo }

func NewIdentityRepo(db *gorm.DB) *IdentityRepo { return &IdentityRepo{repo.BaseRepo{Connection: db}} }
func (r *IdentityRepo) FindOrCreate(ctx context.Context, appID, openID string) (*model.User, error) {
	now := time.Now().UTC()
	user := model.User{WeChatAppID: appID, WeChatOpenID: openID, Nickname: "微信用户", Status: 1, CreateTime: now, UpdateTime: now, CreateUserID: 0, UpdateUserID: 0}
	// Permanent app/openid uniqueness prevents soft-deleted identities re-registering.
	err := r.DB(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "wechat_app_id"}, {Name: "wechat_open_id"}}, DoNothing: true}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	err = r.DB(ctx).Where("wechat_app_id = ? AND wechat_open_id = ?", appID, openID).Take(&user).Error
	return &user, err
}
func (r *IdentityRepo) ByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.DB(ctx).Where("id = ?", id).Take(&user).Error
	return &user, err
}
