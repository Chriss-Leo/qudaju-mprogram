package api

// Identity external APIs: /v1 is added once by the server.
const (
	WeChatLogin  = "/auth/wechat/login"
	AuthMe       = "/auth/me"
	AuthLogout   = "/auth/logout"
	AuthProfile  = "/auth/me/profile"
	Avatar       = "/media/avatars/:name"
	AvatarPrefix = "/v1/media/avatars/"
	HealthLive   = "/health/live"
	HealthReady  = "/health/ready"
)
