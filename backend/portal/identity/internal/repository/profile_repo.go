package repository

import (
	"context"
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/db/repo"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"gorm.io/gorm"
	"time"
)

type ProfileRepo struct{ repo.BaseRepo }

func NewProfileRepo(db *gorm.DB) *ProfileRepo { return &ProfileRepo{repo.BaseRepo{Connection: db}} }
func (r *ProfileRepo) Update(ctx context.Context, id int64, nickname, avatar string) (*model.User, error) {
	result := r.DB(ctx).Model(&model.User{}).Where("id = ? AND status = ?", id, 1).Updates(map[string]interface{}{"nickname": nickname, "avatar_url": avatar, "update_user_id": id, "update_time": time.Now().UTC()})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected != 1 {
		return nil, errors.New("account unavailable")
	}
	var user model.User
	err := r.DB(ctx).Where("id = ? AND status = ?", id, 1).Take(&user).Error
	return &user, err
}
