package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/api"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"image"
	"image/png"
	"os"
	"strings"
	"testing"
)

type testProfiles struct {
	users *testUsers
	fail  bool
}

func (r *testProfiles) Update(ctx context.Context, id int64, name, avatar string) (*model.User, error) {
	if r.fail || ctx.Err() != nil {
		return nil, errors.New("update failed")
	}
	if id != r.users.user.ID {
		return nil, errors.New("wrong user")
	}
	r.users.user.Nickname = name
	r.users.user.AvatarURL = avatar
	return &r.users.user, nil
}
func avatarPNG() []byte {
	var b bytes.Buffer
	_ = png.Encode(&b, image.NewRGBA(image.Rect(0, 0, 32, 24)))
	return b.Bytes()
}
func profileFixture(t *testing.T) (*ProfileSvc, *testProfiles) {
	t.Helper()
	users := &testUsers{user: model.User{ID: 42, Status: 1, Nickname: "微信用户"}}
	profiles := &testProfiles{users: users}
	return &ProfileSvc{Users: users, Profiles: profiles, AvatarDir: t.TempDir()}, profiles
}
func TestProfileRoundTrip(t *testing.T) {
	s, r := profileFixture(t)
	user, err := s.Update(context.Background(), 42, "  小搭子  ", bytes.NewReader(avatarPNG()))
	if err != nil || user.Nickname != "小搭子" || user.ID != "42" || !strings.HasPrefix(user.AvatarURL, api.AvatarPrefix) {
		t.Fatal("profile save failed")
	}
	file, err := s.OpenAvatar(strings.TrimPrefix(user.AvatarURL, api.AvatarPrefix))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	config, kind, err := image.DecodeConfig(file)
	if err != nil || kind != "jpeg" || config.Width != 24 || config.Height != 24 {
		t.Fatal("avatar not normalized")
	}
	firstURL := user.AvatarURL
	user, err = s.Update(context.Background(), 42, "新昵称", nil)
	if err != nil || user.AvatarURL != firstURL || r.users.user.Nickname != "新昵称" {
		t.Fatal("nickname-only edit must retain avatar")
	}
}
func TestProfileRejectsInvalidInput(t *testing.T) {
	s, _ := profileFixture(t)
	for _, name := range []string{"", "a", "微信用户", "hello\nworld", strings.Repeat("搭", 21)} {
		if _, err := s.Update(context.Background(), 42, name, bytes.NewReader(avatarPNG())); err == nil {
			t.Fatal("invalid nickname accepted")
		}
	}
	if _, err := s.Update(context.Background(), 42, "小搭子", nil); err == nil {
		t.Fatal("first profile requires avatar")
	}
	if _, err := s.Update(context.Background(), 99, "小搭子", bytes.NewReader(avatarPNG())); err == nil {
		t.Fatal("unknown account accepted")
	}
	for _, raw := range [][]byte{[]byte("<svg></svg>"), bytes.Repeat([]byte{'a'}, (2<<20)+1)} {
		if _, err := NormalizeAvatar(bytes.NewReader(raw)); err == nil {
			t.Fatal("invalid avatar accepted")
		}
	}
	for _, name := range []string{"../config/config.local.yaml", "avatar.svg", "/etc/passwd"} {
		if file, err := s.OpenAvatar(name); err == nil {
			file.Close()
			t.Fatal("unsafe avatar name accepted")
		}
	}
}
func TestProfileFailureCleansNewFile(t *testing.T) {
	s, r := profileFixture(t)
	r.fail = true
	if _, err := s.Update(context.Background(), 42, "小搭子", bytes.NewReader(avatarPNG())); err == nil {
		t.Fatal("failed database write accepted")
	}
	files, err := os.ReadDir(s.AvatarDir)
	if err != nil || len(files) != 0 {
		t.Fatal("orphaned upload after failed write")
	}
}
