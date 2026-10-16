package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/errs"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model/dto"
	"regexp"
	"strconv"
	"time"
)

type UserStore interface {
	FindOrCreatePhone(context.Context, string, string, string) (*model.User, error)
	ByID(context.Context, int64) (*model.User, error)
}
type SessionStore interface {
	Save(context.Context, string, int64, time.Duration) error
	UserID(context.Context, string) (int64, error)
	Delete(context.Context, string) error
	ClaimCode(context.Context, string) (bool, error)
	AllowLogin(context.Context, string) (bool, error)
}
type WeChat interface {
	Exchange(context.Context, string) (string, error)
	Phone(context.Context, string) (string, error)
}
type IdentitySvc struct {
	Users    UserStore
	Sessions SessionStore
	WeChat   WeChat
	AppID    string
	TTL      time.Duration
}

func NewIdentitySvc(users UserStore, sessions SessionStore, wechat WeChat, appID string, ttl time.Duration) *IdentitySvc {
	return &IdentitySvc{users, sessions, wechat, appID, ttl}
}
func safe(code string) error { return errs.NewBizError(nil, code) }
func publicUser(user *model.User) dto.User {
	return dto.User{ID: strconv.FormatInt(user.ID, 10), Nickname: user.Nickname, AvatarURL: user.AvatarURL}
}
func (s *IdentitySvc) Login(ctx context.Context, code, phoneCode string) (*dto.Login, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if code == "" || len(code) > 256 || phoneCode == "" || len(phoneCode) > 256 {
		return nil, safe(errcode.ParameterError)
	}
	claimed, err := s.Sessions.ClaimCode(ctx, code)
	if err != nil {
		return nil, safe(errcode.SystemBusy)
	}
	if !claimed {
		return nil, safe(errcode.WeChatLoginFailed)
	}
	openID, err := s.WeChat.Exchange(ctx, code)
	if err != nil || openID == "" {
		return nil, safe(errcode.WeChatLoginFailed)
	}
	claimed, err = s.Sessions.ClaimCode(ctx, "phone:"+phoneCode)
	if err != nil {
		return nil, safe(errcode.SystemBusy)
	}
	if !claimed {
		return nil, safe(errcode.WeChatPhoneFailed)
	}
	phone, err := s.WeChat.Phone(ctx, phoneCode)
	if err != nil || !verifiedPhoneShape.MatchString(phone) {
		return nil, safe(errcode.WeChatPhoneFailed)
	}
	user, err := s.Users.FindOrCreatePhone(ctx, s.AppID, openID, phone)
	if err != nil {
		return nil, safe(errcode.SystemBusy)
	}
	if user == nil || user.ID <= 0 || user.Status != 1 || user.DeletedAt.Valid {
		return nil, safe(errcode.AccountDisabled)
	}
	bytes := make([]byte, 32)
	if _, err = rand.Read(bytes); err != nil {
		return nil, safe(errcode.SystemBusy)
	}
	token := base64.RawURLEncoding.EncodeToString(bytes)
	if err = s.Sessions.Save(ctx, token, user.ID, s.TTL); err != nil {
		return nil, safe(errcode.SystemBusy)
	}
	return &dto.Login{AccessToken: token, TokenType: "Bearer", ExpiresIn: int64(s.TTL / time.Second), User: publicUser(user)}, nil
}

var tokenShape = regexp.MustCompile(`^[A-Za-z0-9_-]{43}$`)
var verifiedPhoneShape = regexp.MustCompile(`^\+[1-9][0-9]{6,14}$`)

func (s *IdentitySvc) Me(ctx context.Context, token string) (*dto.User, error) {
	if !tokenShape.MatchString(token) {
		return nil, safe(errcode.NoAuth)
	}
	id, err := s.Sessions.UserID(ctx, token)
	if err != nil {
		return nil, safe(errcode.NoAuth)
	}
	user, err := s.Users.ByID(ctx, id)
	if err != nil || user == nil || user.Status != 1 || user.DeletedAt.Valid {
		return nil, safe(errcode.NoAuth)
	}
	value := publicUser(user)
	return &value, nil
}
func (s *IdentitySvc) Logout(ctx context.Context, token string) error {
	if !tokenShape.MatchString(token) {
		return safe(errcode.NoAuth)
	}
	if err := s.Sessions.Delete(ctx, token); err != nil {
		return safe(errcode.SystemBusy)
	}
	return nil
}
func (s *IdentitySvc) AllowLogin(ctx context.Context, ip string) bool {
	allowed, err := s.Sessions.AllowLogin(ctx, ip)
	return err == nil && allowed
}
