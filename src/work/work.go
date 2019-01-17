package work

import (
	"appletMessagesServer/src/cron"
	"appletMessagesServer/src/http"
	"appletMessagesServer/src/redispool"
	"log"
)

func Start() {
	go cron.Start()
	go http.Start()
	log.Println("server start ok")
}

func Stop() {
	cron.Stop()
	http.Stop()
	redispool.Local.Close()
	log.Println("server stop ok")
}
