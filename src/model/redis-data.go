package model

import (
	"appletMessagesServer/src/redispool"

	"github.com/garyburd/redigo/redis"
)

//GetRedisAccessToken 从redis中获取缓存的token
func GetRedisAccessToken(key string) string {
	rc := redispool.Local.Get()
	defer rc.Close()
	accessToken, _ := redis.String(rc.Do("HGET", key, "token"))
	return accessToken
}
