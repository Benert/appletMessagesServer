package g

import (
	"appletMessagesServer/src/logger"
	"appletMessagesServer/src/redispool"
	"appletMessagesServer/src/util"

	"encoding/json"
	"io/ioutil"
	"sync"
)

type (
	AppletConfig struct {
		AppName      string `json:"appName"`
		AppId        string `json:"appId"`
		AppSecret    string `json:"appSecret"`
		Token        string `json:"token"`
		Aeskey       string `json:"aeskey"`
		Debug        bool   `json:"debug"`
		ActionAddres string `josn:"actionAddres"`
		AutoAnswer   bool   `json:"autoAnswer"`
		Welcome      string `json:"welcome"`
	}
	Http struct {
		Enable bool   `json:"enable"`
		Listen string `json:"listen"`
	}
	Config struct {
		Applet []*AppletConfig  `json:"applet"`
		Log    logger.Config    `json:"log"`
		Redis  redispool.Config `json:"redis"`
		HTTP   Http             `json:"http"`
	}
)

var (
	Root       string
	configLock = new(sync.RWMutex)
	ConfigData *Config
)

//Get 解析配置
func GetConf(v interface{}) {
	bs, err := ioutil.ReadFile("cfg.json")
	util.FailOnError(err, "error opening configfile")
	err = json.Unmarshal(bs, v)
	util.FailOnError(err, "EROR: wrong config json")
}

//Init 初始化配置文件
func InitConf(c Config) {
	logger.Init(c.Log)
	redispool.Init(c.Redis)
}
