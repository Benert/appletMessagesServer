package cron

import (
	"github.com/imroc/log"
)

//Start 定时任务启动
func Start() {
	StartToken() //获取accesstoken
	log.Info("cron.Start ok")
}

//Stop 定时任务结束
func Stop() {
	log.Info("cron.Stop ok")
}
