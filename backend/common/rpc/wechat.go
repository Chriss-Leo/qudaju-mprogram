package rpc

import (
	"context"
	"errors"
	"resty.dev/v3"
	"time"
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
}

func NewWeChatClient(appID, secret string) *WeChatClient {
	return &WeChatClient{HTTP: resty.New().SetTimeout(8 * time.Second), AppID: appID, AppSecret: secret, endpoint: "https://api.weixin.qq.com/sns/jscode2session"}
}
func (c *WeChatClient) Exchange(ctx context.Context, code string) (string, error) {
	var data WeChatIdentity
	response, err := c.HTTP.R().SetContext(ctx).SetQueryParams(map[string]string{"appid": c.AppID, "secret": c.AppSecret, "js_code": code, "grant_type": "authorization_code"}).SetResult(&data).Get(c.endpoint)
	// Never propagate Resty's error/URL/body: they may contain credentials or session_key.
	if err != nil || response == nil || response.StatusCode() != 200 || data.ErrCode != 0 || data.OpenID == "" || data.SessionKey == "" {
		return "", errors.New("wechat exchange failed")
	}
	return data.OpenID, nil
}
