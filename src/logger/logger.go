package logger

import "github.com/imroc/log"

type Config struct {
	Debug  bool   `json:"debug"`
	Output string `json:"output"`
}

func Init(c Config) {
	log.SetDebug(c.Debug)
	if c.Output != "" && c.Output != "std" {
		err := log.SetFilename(c.Output)
		if err != nil {
			panic(err)
		}
	}
}
