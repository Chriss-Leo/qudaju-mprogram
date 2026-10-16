package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/api"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/request/middleware"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/service"
	"github.com/gin-gonic/gin"
	"image"
	"image/png"
	"mime/multipart"
	"net/http/httptest"
	"strings"
	"testing"
)

type profileFixture struct{ user model.User }

func (r *profileFixture) ByID(ctx context.Context, id int64) (*model.User, error) {
	if id != r.user.ID {
		return nil, errors.New("wrong account")
	}
	return &r.user, nil
}
func (r *profileFixture) FindOrCreatePhone(context.Context, string, string, string) (*model.User, error) {
	return nil, errors.New("not used")
}
func (r *profileFixture) Update(ctx context.Context, id int64, name, url string) (*model.User, error) {
	if id != 42 {
		return nil, errors.New("wrong account")
	}
	r.user.Nickname = name
	r.user.AvatarURL = url
	return &r.user, nil
}
func TestProfileHTTPUploadAndRead(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := &profileFixture{user: model.User{ID: 42, Status: 1, Nickname: "微信用户"}}
	svc := &service.ProfileSvc{Users: repo, Profiles: repo, AvatarDir: t.TempDir()}
	router := gin.New()
	middleware.RegisterMdl(router, nil)
	auth := middleware.RequireLogin(func(ctx context.Context, token string) (middleware.Principal, error) {
		if token != "test-token" {
			return middleware.Principal{}, errors.New("unauthorized")
		}
		return middleware.Principal{UserID: "42", RoleCodes: []string{"member"}}, nil
	})
	router.POST("/v1"+api.AuthProfile, auth, middleware.ApiPermission("member"), UpdateProfile(svc))
	router.GET("/v1"+api.Avatar, Avatar(svc))
	denied := httptest.NewRecorder()
	router.ServeHTTP(denied, httptest.NewRequest("POST", "/v1"+api.AuthProfile, nil))
	if denied.Code != 401 {
		t.Fatal("unauthenticated profile update allowed")
	}
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	_ = writer.WriteField("nickname", "小搭子")
	_ = writer.WriteField("userId", "99")
	file, _ := writer.CreateFormFile("avatar", "photo.png")
	_ = png.Encode(file, image.NewRGBA(image.Rect(0, 0, 8, 8)))
	_ = writer.Close()
	req := httptest.NewRequest("POST", "/v1"+api.AuthProfile, &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer test-token")
	result := httptest.NewRecorder()
	router.ServeHTTP(result, req)
	var payload struct {
		Code string                                   `json:"code"`
		Data struct{ ID, Nickname, AvatarURL string } `json:"data"`
	}
	if json.Unmarshal(result.Body.Bytes(), &payload) != nil || payload.Code != "1000000" || payload.Data.ID != "42" || payload.Data.Nickname != "小搭子" {
		t.Fatal("upload contract failed")
	}
	avatar := httptest.NewRecorder()
	router.ServeHTTP(avatar, httptest.NewRequest("GET", payload.Data.AvatarURL, nil))
	if avatar.Code != 200 || avatar.Header().Get("Content-Type") != "image/jpeg" {
		t.Fatal("stored avatar unavailable")
	}
	edit := httptest.NewRequest("POST", "/v1"+api.AuthProfile, strings.NewReader("nickname=Other"))
	edit.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	edit.Header.Set("Authorization", "Bearer test-token")
	edited := httptest.NewRecorder()
	router.ServeHTTP(edited, edit)
	if repo.user.Nickname != "Other" {
		t.Fatal("nickname edit failed")
	}
	oversized := httptest.NewRequest("POST", "/v1"+api.AuthProfile, strings.NewReader(strings.Repeat("x", (3<<20)+1)))
	oversized.Header.Set("Content-Type", "multipart/form-data; boundary=anything")
	oversized.Header.Set("Authorization", "Bearer test-token")
	rejected := httptest.NewRecorder()
	router.ServeHTTP(rejected, oversized)
	if strings.Contains(rejected.Body.String(), `"code":"1000000"`) {
		t.Fatal("oversized upload accepted")
	}
}
