package cache

import "github.com/go-redis/redis/v8"

// RDB 由目标项目启动过程注入；本模板不会自动连接 Redis。
var RDB redis.UniversalClient
