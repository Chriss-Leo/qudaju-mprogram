# Go 后端

按 `go-microservice-standard` 建立 Go 多模块工程：共享 `common` 和 `portal/identity` 身份服务。使用 Gin、GORM/PostgreSQL、Redis v8、Viper、Logrus、Resty；显式装配依赖，controller → service → repository。完整响应 helper 沿用规范资产，不在 controller 自定义业务响应。

## 启动

需要 Go 1.25、PostgreSQL、Redis，以及与小程序一致的微信 AppID / AppSecret。这里不包含数据库或 Redis 的自动安装。先在目标数据库执行 `script/001_identity.up.sql`，不会启动时自动修改表结构。

从 `backend/portal/identity` 目录运行。默认读取 `config/config.yaml`；配置值可通过以下环境变量覆盖：

| 变量 | 用途 |
| --- | --- |
| APP_POSTGRES_DSN | PostgreSQL DSN，包含数据库与凭据 |
| APP_REDIS_ADDR / APP_REDIS_PASSWORD / APP_REDIS_DB | Redis 连接配置 |
| APP_WECHAT_APP_ID / APP_WECHAT_APP_SECRET | 微信服务端登录凭据 |
| APP_HTTP_PORT | HTTP 端口，默认 8081 |
| APP_SESSION_TTL | 会话有效期，默认 2h，允许 1m–24h |
| APP_LOG_LEVEL | 日志等级，默认 info |

`APP_ENV=local` 时读取 `config/config.local.yaml`，此文件已忽略。不要提交真实凭据；生产使用部署环境注入。H5 联调需要在 YAML `server.origins` 中明确允许前端 origin，如 `http://127.0.0.1:5173`，不使用通配符。

```bash
cd backend/portal/identity
GOWORK=off go mod download
GOWORK=off go run .
```

缺少配置、迁移、数据库、Redis 或语言资源时启动失败。`/health/live` 检查进程，`/health/ready` 检查 PostgreSQL 与 Redis；编译通过不代表真实微信联调完成。

## 登录接口契约

- `POST /v1/auth/wechat/login`：JSON `{"code":"uni.login 返回的一次性 code"}`。
- `GET /v1/auth/me`：`Authorization: Bearer <accessToken>`。
- `POST /v1/auth/logout`：同上，撤销当前会话。

成功返回 `{"code":"1000000","msg":"…","data":…, "time":"Unix 秒"}`。登录 data 为 `{accessToken,tokenType:"Bearer",expiresIn,user:{id,nickname,avatarUrl}}`；`expiresIn` 单位秒，用户 ID 为十进制字符串。普通业务失败仍为 HTTP 200，请检查业务 code；身份失效为 HTTP 401。前端不得凭 code 自行创建登录状态。

微信交换仅服务端调用 `jscode2session`，不会返回 openid、session_key 或 AppSecret。32 字节随机不透明 token 的 SHA-256 摘要保存在 Redis；支持过期与注销，账号禁用/删除检查在请求时执行。一次性 code 防重复、每 IP 每分钟最多 10 次登录，Redis 故障时拒绝登录。日志不记录 token、code、凭据或微信请求 URL。

生产接口必须使用 HTTPS，并在微信后台配置 request 合法域名。微信登录不自动获取头像、昵称、手机号；当前账户返回默认资料，资料编辑需后续独立接口。

## 工程边界与测试

当前服务仅实现身份接口；组局发布、审核、名额、基金支付、聊天与推送的服务端接口未在此次接入。前端已有本地流程不能当作服务端授权。普通活动用户只具有 member 角色；未来管理角色/菜单权限、OAuth2 授权服务及内部服务接口需要独立实现，不允许客户端请求头绕过鉴权。

分别在两个模块运行检查（backend 根目录不是模块，不能使用 `GOWORK=off go test ./...`）：

```bash
cd backend/common
GOWORK=off go test -race ./...
GOWORK=off go vet ./...
# 另一个终端，项目根目录开始
cd backend/portal/identity
GOWORK=off go test -race ./...
GOWORK=off go vet ./...
```

测试覆盖微信失败、code 并发重放、会话撤销、账户禁用、权限拦截和响应格式。使用注入替身的单元测试不代替真实 PostgreSQL/Redis/微信端到端验收。
