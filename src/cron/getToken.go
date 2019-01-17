package cron

import (
	"appletMessagesServer/src/g"
	"appletMessagesServer/src/model"
	redispool "appletMessagesServer/src/redispool"
	"appletMessagesServer/src/util"

	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/imroc/log"
)

//monitorToken 定时检查 token，如果发现要过期了，重现请求一个
func monitorToken() {
	log.Info("monitorToken start")
	for {
		rc := redispool.Local.Get()
		AppletArr := g.ConfigTo().Applet
		for _, c := range AppletArr {
			appid := c.AppId
			appSecret := c.AppSecret
			key := "applet_token_" + appid // 判断 redis中这个token 是否存
			t, _ := redis.Int64(rc.Do("TTL", key))

			if c.Debug == true {
				log.Info("key: ", key, "  appid: ", appid, " secret: ", appSecret, " t:", t, "\n token: ", model.GetRedisAccessToken(key))
			}

			if t < 600 { //  即将过期

				token := util.GetToken(appid, appSecret)
				if token == nil {
					continue
				}
				log.Info("wx access token refresh", "***"+token.Token[12:20]+"***", token.ExpiresIn)
				rc.Do("HMSET", key, "token", token.Token)
				rc.Do("EXPIRE", key, token.ExpiresIn-100) // 留一个保护间隔
			}
		}

		rc.Close()

		time.Sleep(30 * time.Second) // 定时检查 token
	}
}

//StartToken 监控token的有效性
func StartToken() {
	go monitorToken()
}
