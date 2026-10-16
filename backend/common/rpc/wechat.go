package rpc

import (
	"context"
	"errors"
	"regexp"
	"sync"
	"time"

	"resty.dev/v3"
)

type WeChatIdentity struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
}
type WeChatClient struct {
	HTTP             *resty.Client
	AppID, AppSecret string
	endpoint         string
	phoneEndpoint    string
	tokenEndpoint    string
	tokenMu          sync.Mutex
	accessToken      string
	tokenExpires     time.Time
}

func NewWeChatClient(appID, secret string) *WeChatClient {
	return &WeChatClient{HTTP: resty.New().SetTimeout(8 * time.Second), AppID: appID, AppSecret: secret, endpoint: "https://api.weixin.qq.com/sns/jscode2session", phoneEndpoint: "https://api.weixin.qq.com/wxa/business/getuserphonenumber", tokenEndpoint: "https://api.weixin.qq.com/cgi-bin/stable_token"}
}

// Stable access tokens stay server-side. The stable-token API does not invalidate
// another instance's token; this process cache expires early and coalesces refreshes.
func (c *WeChatClient) serverToken(ctx context.Context, rejected string) (string, error) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	if c.accessToken != "" && c.accessToken != rejected && time.Now().Before(c.tokenExpires) {
		return c.accessToken, nil
	}
	var data struct {
		Token   string `json:"access_token"`
		Expires int    `json:"expires_in"`
		ErrCode int    `json:"errcode"`
	}
	response, err := c.HTTP.R().SetContext(ctx).SetBody(map[string]any{"grant_type": "client_credential", "appid": c.AppID, "secret": c.AppSecret, "force_refresh": rejected != ""}).SetForceResponseContentType("application/json").SetResult(&data).Post(c.tokenEndpoint)
	if err != nil || response == nil || response.StatusCode() != 200 || data.ErrCode != 0 || data.Token == "" || data.Expires <= 60 || data.Expires > 86400 {
		return "", errors.New("wechat server token unavailable")
	}
	c.accessToken, c.tokenExpires = data.Token, time.Now().Add(time.Duration(data.Expires-60)*time.Second)
	return c.accessToken, nil
}

var phoneShape = regexp.MustCompile(`^[1-9][0-9]{5,14}$`)
var countryShape = regexp.MustCompile(`^[1-9][0-9]{0,3}$`)

// Phone exchanges the one-use native getPhoneNumber code, never client phone text.
func (c *WeChatClient) Phone(ctx context.Context, code string) (string, error) {
	rejected := ""
	for attempt := 0; attempt < 2; attempt++ {
		token, err := c.serverToken(ctx, rejected)
		if err != nil {
			return "", err
		}
		var data struct {
			ErrCode int `json:"errcode"`
			Info    struct {
				Pure      string `json:"purePhoneNumber"`
				Country   string `json:"countryCode"`
				Watermark struct {
					AppID string `json:"appid"`
				} `json:"watermark"`
			} `json:"phone_info"`
		}
		response, err := c.HTTP.R().SetContext(ctx).SetQueryParam("access_token", token).SetBody(map[string]string{"code": code}).SetForceResponseContentType("application/json").SetResult(&data).Post(c.phoneEndpoint)
		if err != nil || response == nil || response.StatusCode() != 200 {
			return "", errors.New("wechat phone exchange failed")
		}
		if attempt == 0 && (data.ErrCode == 40001 || data.ErrCode == 40014 || data.ErrCode == 42001) {
			rejected = token
			continue
		}
		if data.ErrCode != 0 || data.Info.Watermark.AppID != c.AppID || !countryShape.MatchString(data.Info.Country) || !phoneShape.MatchString(data.Info.Pure) || len(data.Info.Country+data.Info.Pure) > 15 {
			return "", errors.New("wechat phone authorization failed")
		}
		return "+" + data.Info.Country + data.Info.Pure, nil
	}
	return "", errors.New("wechat phone exchange failed")
}
func (c *WeChatClient) Exchange(ctx context.Context, code string) (string, error) {
	var data WeChatIdentity
	// This fixed endpoint returns JSON; do not depend on the upstream MIME type.
	response, err := c.HTTP.R().
		SetContext(ctx).
		SetQueryParams(map[string]string{"appid": c.AppID, "secret": c.AppSecret, "js_code": code, "grant_type": "authorization_code"}).
		SetForceResponseContentType("application/json").
		SetResult(&data).
		Get(c.endpoint)
	// Never propagate Resty's error/URL/body: they may contain credentials or session_key.
	if err != nil || response == nil || response.StatusCode() != 200 || data.ErrCode != 0 || data.OpenID == "" || data.SessionKey == "" {
		return "", errors.New("wechat exchange failed")
	}
	return data.OpenID, nil
}
