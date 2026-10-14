package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type SessionRepo struct{ Client redis.UniversalClient }

func NewSessionRepo(client redis.UniversalClient) *SessionRepo { return &SessionRepo{Client: client} }
func hash(value string) string                                 { h := sha256.Sum256([]byte(value)); return hex.EncodeToString(h[:]) }
func sessionKey(token string) string                           { return "qudaju:identity:session:" + hash(token) }
func (r *SessionRepo) Save(ctx context.Context, token string, userID int64, ttl time.Duration) error {
	return r.Client.Set(ctx, sessionKey(token), strconv.FormatInt(userID, 10), ttl).Err()
}
func (r *SessionRepo) UserID(ctx context.Context, token string) (int64, error) {
	value, err := r.Client.Get(ctx, sessionKey(token)).Result()
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid session")
	}
	return id, nil
}
func (r *SessionRepo) Delete(ctx context.Context, token string) error {
	return r.Client.Del(ctx, sessionKey(token)).Err()
}
func (r *SessionRepo) ClaimCode(ctx context.Context, code string) (bool, error) {
	return r.Client.SetNX(ctx, "qudaju:identity:code:"+hash(code), "1", 5*time.Minute).Result()
}

var rateScript = redis.NewScript(`local n=redis.call('INCR',KEYS[1]); if n==1 then redis.call('EXPIRE',KEYS[1],60) end; return n`)

func (r *SessionRepo) AllowLogin(ctx context.Context, ip string) (bool, error) {
	n, err := rateScript.Run(ctx, r.Client, []string{"qudaju:identity:rate:" + hash(ip)}).Int64()
	return n <= 10, err
}
