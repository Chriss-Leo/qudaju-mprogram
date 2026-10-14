package api

// Identity external APIs: /v1 is added once by the server.
const (
	WeChatLogin = "/auth/wechat/login"
	AuthMe      = "/auth/me"
	AuthLogout  = "/auth/logout"
	HealthLive  = "/health/live"
	HealthReady = "/health/ready"
)
