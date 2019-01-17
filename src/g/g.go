package g

import (
	"appletMessagesServer/src/util"

	"os"
	"runtime"

	"github.com/imroc/log"
)

const (
	VERSION = "1.0.0"
)

//InitRootDir 设置文件root
func InitRootDir() {
	var err error
	Root, err = os.Getwd()
	util.CheckErrors(err, "getwd fail:")
	log.Info("Root:", Root)
}

//Init 初始化配置
func Init(c Config) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	InitRootDir()
	InitConf(c)
	ConfigData = &c
}

func ConfigTo() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return ConfigData
}
