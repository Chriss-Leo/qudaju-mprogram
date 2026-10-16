package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/api"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/common/consts/errcode"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model/dto"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ProfileStore interface {
	Update(context.Context, int64, string, string) (*model.User, error)
}
type ProfileSvc struct {
	Users     UserStore
	Profiles  ProfileStore
	AvatarDir string
}

func ValidNickname(value string) bool {
	n := utf8.RuneCountInString(value)
	if n < 2 || n > 20 || value == "微信用户" {
		return false
	}
	for _, r := range value {
		if unicode.IsControl(r) {
			return false
		}
	}
	return true
}

// Normalize rejects disguised/non-image uploads and bounds decode memory and output size.
func NormalizeAvatar(source io.Reader) ([]byte, error) {
	raw, err := io.ReadAll(io.LimitReader(source, (2<<20)+1))
	if err != nil || len(raw) > 2<<20 {
		return nil, safe(errcode.ParameterError)
	}
	config, format, err := image.DecodeConfig(bytes.NewReader(raw))
	if err != nil || (format != "jpeg" && format != "png") || config.Width < 1 || config.Height < 1 || config.Width > 4096 || config.Height > 4096 || int64(config.Width)*int64(config.Height) > 16000000 {
		return nil, safe(errcode.ParameterError)
	}
	src, _, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		return nil, safe(errcode.ParameterError)
	}
	size := config.Width
	if config.Height < size {
		size = config.Height
	}
	target := size
	if target > 512 {
		target = 512
	}
	dst := image.NewRGBA(image.Rect(0, 0, target, target))
	left, top := (config.Width-size)/2, (config.Height-size)/2
	for y := 0; y < target; y++ {
		for x := 0; x < target; x++ {
			dst.Set(x, y, src.At(src.Bounds().Min.X+left+x*size/target, src.Bounds().Min.Y+top+y*size/target))
		}
	}
	var out bytes.Buffer
	if jpeg.Encode(&out, dst, &jpeg.Options{Quality: 85}) != nil {
		return nil, safe(errcode.SystemBusy)
	}
	return out.Bytes(), nil
}
func (s *ProfileSvc) Update(ctx context.Context, id int64, nickname string, avatar io.Reader) (*dto.User, error) {
	nickname = strings.TrimSpace(nickname)
	if !ValidNickname(nickname) {
		return nil, safe(errcode.ParameterError)
	}
	user, err := s.Users.ByID(ctx, id)
	if err != nil || user == nil || user.Status != 1 || user.DeletedAt.Valid {
		return nil, safe(errcode.NoAuth)
	}
	path := ""
	url := user.AvatarURL
	if avatar != nil {
		data, err := NormalizeAvatar(avatar)
		if err != nil {
			return nil, err
		}
		if ctx.Err() != nil {
			return nil, safe(errcode.SystemBusy)
		}
		nonce := make([]byte, 24)
		if _, err = rand.Read(nonce); err != nil {
			return nil, safe(errcode.SystemBusy)
		}
		name := hex.EncodeToString(nonce) + ".jpg"
		path = filepath.Join(s.AvatarDir, name)
		if err = os.MkdirAll(s.AvatarDir, 0700); err != nil {
			return nil, safe(errcode.SystemBusy)
		}
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			return nil, safe(errcode.SystemBusy)
		}
		_, writeErr := file.Write(data)
		closeErr := file.Close()
		if writeErr != nil || closeErr != nil {
			_ = os.Remove(path)
			return nil, safe(errcode.SystemBusy)
		}
		url = api.AvatarPrefix + name
	}
	if url == "" {
		return nil, safe(errcode.ParameterError)
	}
	updated, err := s.Profiles.Update(ctx, id, nickname, url)
	if err != nil {
		if path != "" {
			_ = os.Remove(path)
		}
		return nil, safe(errcode.SystemBusy)
	}
	value := publicUser(updated)
	return &value, nil
}

var avatarName = regexp.MustCompile(`^[a-f0-9]{48}\.jpg$`)

func (s *ProfileSvc) OpenAvatar(name string) (*os.File, error) {
	if !avatarName.MatchString(name) {
		return nil, os.ErrNotExist
	}
	return os.Open(filepath.Join(s.AvatarDir, name))
}
