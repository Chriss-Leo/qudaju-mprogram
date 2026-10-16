# Go 后端

按 `go-microservice-standard` 建立 Go 多模块工程：共享 `common` 和 `portal/identity` 身份服务。使用 Gin、GORM/PostgreSQL、Redis v8、Viper、Logrus、Resty；显式装配依赖，controller → service → repository。完整响应 helper 沿用规范资产，不在 controller 自定义业务响应。

## 启动

需要 Go 1.25、PostgreSQL、Redis，以及与小程序一致的微信 AppID / AppSecret。这里不包含数据库或 Redis 的自动安装。新安装在 `qudaju` 数据库执行 `script/pg.sql`，所有表直接创建在 `public`，不会启动时自动修改表结构。旧版两个 schema 的升级使用 `script/002_move_identity_to_public.up.sql`：保留数据迁表后删除空 schema，不使用 CASCADE。`script/legacy` 只保留已执行过的历史记录，不用于新安装。

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

本机现有 Docker 联调配置保存在已忽略的 `config/config.local.yaml`，对应 PostgreSQL `127.0.0.1:5432/qudaju` 和 Redis `127.0.0.1:6379`；密码仅在本地文件中。启动该配置：

```bash
cd backend/portal/identity
APP_ENV=local GOWORK=off go run .
```

前端 `.env.local` 指向 `http://127.0.0.1:8081`，仅限电脑开发模式；需重启前端开发编译。手机预览不能使用该回环地址，生产构建也不能使用 HTTP。

### VS Code 调试启动

用 VS Code 打开整个 `qudaju` 项目根目录，安装推荐的 Go 扩展（`golang.go`）。进入「运行和调试」，选择「趣搭局后端 · identity（本地）」后按 F5；首次提示安装 Delve 调试器时按扩展提示安装。

`.vscode/launch.json` 已设置服务工作目录、`APP_ENV=local` 与 `GOWORK=off`，读取 `config/config.local.yaml`，不在调试配置中保存密钥。先保持 PostgreSQL / Redis 容器运行，并完成数据库迁移。默认监听 8081，停止按钮会结束本次调试进程；不要同时用终端启动另一份占用同一端口的服务。

缺少配置、迁移、数据库、Redis 或语言资源时启动失败。`/health/live` 检查进程，`/health/ready` 检查 PostgreSQL 与 Redis；编译通过不代表真实微信联调完成。

## 登录接口契约

### 微信原生手机号注册 / 登录

登录按钮使用 `open-type="getPhoneNumber"`，仅微信原生授权弹窗，不再在登录后显示自定义头像昵称弹层。拒绝或取消授权不会请求注册接口；H5仅提示在微信小程序中操作，不模拟成功。头像昵称可稍后通过「我的 → 编辑资料」自愿设置。

客户端获取两个不同的一次性凭证：`uni.login` 的 `code` 用于获取 OpenID；原生 `getphonenumber` 事件的 `detail.code` 作为 `phoneCode` 换取手机号。登录接口二者均为必填，不接受手机号字符串替代授权。Go 在两次微信验证都成功后才创建账号或更新已验证手机号，并签发 Redis 会话；失败时不创建账号、不发会话。

现有数据库先执行 `script/003_wechat_phone.up.sql`，为 `public.auth_user` 增加可空 `phone_number` 和 `phone_verified_time`；历史用户及其 ID、昵称、头像保留，下次手机号登录时补全。新装仍用 `pg.sql`。账号以 AppID + OpenID 唯一关联，不因手机号相同自动合并账号；手机号不通过 user DTO、日志或公开信息卡片返回。

服务端使用 `stable_token` 获取应用访问令牌，进程内缓存至提前60秒到期并合并并发刷新；微信明确返回令牌失效时最多刷新重试一次。此令牌只在后端使用，和用户会话 token 不同。整条登录链路超时15秒；前端等待20秒。密钥、code、令牌及微信原始响应不得写日志。

真实联调需确保微信后台已具备手机号能力及可用额度，隐私保护指引中声明手机号用途，并按微信要求配置服务器 IP 白名单、合法域名。使用真机验证原生授权；浏览器不能弹出微信手机号选择界面。后端变更后请在 VS Code 停止并重新 F5，前端重新编译微信小程序。本实现没有代替你开通微信平台能力。

### 可选头像与昵称编辑

仅「我的 → 编辑资料」显示 `WeChatProfileSheet`。微信使用 `button open-type="chooseAvatar"` 与 `input type="nickname"`；选择头像回调获取 `e.detail.avatarUrl` 临时路径，确认时用 `uni.uploadFile` 将文件和昵称提交当前用户接口，保存成功后更新资料。不是调用旧的用户资料授权接口，也不强制用户在手机号注册时填写。

`POST /v1/auth/me/profile` 需要 Bearer 会话及 member 角色。首次提交使用 multipart/form-data，字段 `nickname` 与文件 `avatar`；已存在头像时可以只用 form-urlencoded 提交 `nickname`。用户 ID 始终从会话取，忽略客户端传入的账号/头像 URL。返回标准成功 envelope 的 user，字段与 `/me` 一致。

昵称去除首尾空白后需2–20个 Unicode 字符、不含控制符，不接受初始占位名「微信用户」。头像只接受2MiB以内的真实 JPG/PNG，校验解码尺寸，中心裁成正方形并重编码为最大512px JPEG，原文件名不使用。公开读取路径 `/v1/media/avatars/<随机文件名>.jpg`；该相对路径由前端按接口 origin 解析，本地 HTTP 仅开发模式允许。头像是公开资料，不上传敏感图片。

文件保存位置由 `media.avatar_dir` 配置，默认服务目录 `./storage/avatars`，已被 Git 忽略；用户表沿用现有 nickname/avatar_url，无需新增迁移。数据库更新失败会移除本次新文件，历史头像不自动清理。部署必须持久化该目录，多实例需共享文件存储或后续迁移对象存储；只备份数据库不能恢复头像。小程序上线还需配置 uploadFile / 图片访问相关合法域名。当前完成格式校验，不等于已接入服务端违规图片/文本审核；生产开放资料前需接入相应内容安全能力。

- `POST /v1/auth/wechat/login`：JSON `{"code":"uni.login 返回的一次性 code","phoneCode":"原生手机号授权返回的一次性 code"}`。
- `GET /v1/auth/me`：`Authorization: Bearer <accessToken>`。
- `POST /v1/auth/logout`：同上，撤销当前会话。

成功返回 `{"code":"1000000","msg":"…","data":…, "time":"Unix 秒"}`。登录 data 为 `{accessToken,tokenType:"Bearer",expiresIn,user:{id,nickname,avatarUrl}}`；`expiresIn` 单位秒，用户 ID 为十进制字符串。普通业务失败仍为 HTTP 200，请检查业务 code；身份失效为 HTTP 401。前端不得凭 code 自行创建登录状态。

微信交换由服务端调用 `jscode2session` 和 `getuserphonenumber`，不会返回 openid、session_key、完整手机号或 AppSecret。32 字节随机不透明 token 的 SHA-256 摘要保存在 Redis；支持过期与注销，账号禁用/删除检查在请求时执行。两种一次性 code 分开防重复、每 IP 每分钟最多 10 次登录，Redis 故障时拒绝登录。日志不记录 token、code、凭据或微信请求 URL。

生产接口必须使用 HTTPS，并在微信后台配置 request 合法域名。手机号仅在用户主动原生授权后获取；头像、昵称不自动读取。

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
