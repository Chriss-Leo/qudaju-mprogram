package service

import (
	"context"
	"errors"
	"github.com/Chriss-Leo/qudaju-mprogram/backend/portal/identity/internal/model"
	"strings"
	"sync"
	"testing"
	"time"
)

type testUsers struct {
	user    model.User
	creates int
}

func (u *testUsers) FindOrCreate(_ context.Context, appID, openID string) (*model.User, error) {
	if appID != "test-app" || openID != "test-openid" {
		return nil, errors.New("unexpected identity")
	}
	u.creates++
	return &u.user, nil
}
func (u *testUsers) ByID(_ context.Context, id int64) (*model.User, error) {
	if id != u.user.ID {
		return nil, errors.New("not found")
	}
	return &u.user, nil
}

type testSessions struct {
	sync.Mutex
	tokens map[string]int64
	codes  map[string]bool
	ttl    time.Duration
	fail   bool
}

func (s *testSessions) Save(_ context.Context, token string, id int64, ttl time.Duration) error {
	s.Lock()
	defer s.Unlock()
	if s.fail {
		return errors.New("redis unavailable")
	}
	s.tokens[token] = id
	s.ttl = ttl
	return nil
}
func (s *testSessions) UserID(_ context.Context, token string) (int64, error) {
	s.Lock()
	defer s.Unlock()
	id, ok := s.tokens[token]
	if !ok || s.fail {
		return 0, errors.New("missing")
	}
	return id, nil
}
func (s *testSessions) Delete(_ context.Context, token string) error {
	s.Lock()
	defer s.Unlock()
	delete(s.tokens, token)
	return nil
}
func (s *testSessions) ClaimCode(_ context.Context, code string) (bool, error) {
	s.Lock()
	defer s.Unlock()
	if s.fail {
		return false, errors.New("redis unavailable")
	}
	if s.codes[code] {
		return false, nil
	}
	s.codes[code] = true
	return true, nil
}
func (s *testSessions) AllowLogin(context.Context, string) (bool, error) { return !s.fail, nil }

type testWeChat struct{ fail bool }

func (w testWeChat) Exchange(context.Context, string) (string, error) {
	if w.fail {
		return "", errors.New("sensitive upstream body")
	}
	return "test-openid", nil
}
func fixture() (*IdentitySvc, *testUsers, *testSessions) {
	users := &testUsers{user: model.User{ID: 9223372036854770000, Nickname: "微信用户", Status: 1}}
	sessions := &testSessions{tokens: map[string]int64{}, codes: map[string]bool{}}
	return NewIdentitySvc(users, sessions, testWeChat{}, "test-app", 2*time.Hour), users, sessions
}
func TestLoginSessionLogout(t *testing.T) {
	svc, users, sessions := fixture()
	ctx := context.Background()
	login, err := svc.Login(ctx, "test-code")
	if err != nil {
		t.Fatal(err)
	}
	if len(login.AccessToken) != 43 || login.ExpiresIn != 7200 || sessions.ttl != 2*time.Hour {
		t.Fatal("invalid token/TTL")
	}
	if login.User.ID != "9223372036854770000" {
		t.Fatal("BIGINT must remain a decimal string")
	}
	if _, err = svc.Login(ctx, "test-code"); err == nil {
		t.Fatal("code replay accepted")
	}
	if users.creates != 1 {
		t.Fatal("replayed identity mutation")
	}
	if _, err = svc.Me(ctx, login.AccessToken); err != nil {
		t.Fatal(err)
	}
	if err = svc.Logout(ctx, login.AccessToken); err != nil {
		t.Fatal(err)
	}
	if _, err = svc.Me(ctx, login.AccessToken); err == nil {
		t.Fatal("revoked token accepted")
	}
}
func TestRejectFailureAndDisabled(t *testing.T) {
	ctx := context.Background()
	svc, users, sessions := fixture()
	svc.WeChat = testWeChat{fail: true}
	_, err := svc.Login(ctx, "failure-code")
	if err == nil || strings.Contains(err.Error(), "sensitive") {
		t.Fatal("failed login leaked cause or succeeded")
	}
	if users.creates != 0 {
		t.Fatal("upstream failure wrote user")
	}
	svc.WeChat = testWeChat{}
	users.user.Status = 0
	if _, err = svc.Login(ctx, "disabled-code"); err == nil || len(sessions.tokens) != 0 {
		t.Fatal("disabled user got session")
	}
	sessions.fail = true
	if _, err = svc.Login(ctx, "redis-code"); err == nil {
		t.Fatal("redis failure opened login")
	}
	if _, err = svc.Me(ctx, strings.Repeat("a", 43)); err == nil {
		t.Fatal("missing token accepted")
	}
}
func TestConcurrentCodeOnlyOnce(t *testing.T) {
	svc, _, sessions := fixture()
	var wg sync.WaitGroup
	var mu sync.Mutex
	success := 0
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := svc.Login(context.Background(), "one-code"); err == nil {
				mu.Lock()
				success++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if success != 1 || len(sessions.tokens) != 1 {
		t.Fatal("concurrent replay issued sessions")
	}
}
