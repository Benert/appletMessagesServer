package redispool

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/imroc/log"
)

var Local *redis.Pool

type Config struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	MaxIdle  int    `json:"maxIdle"`
	Db       int    `json:"db"`
}

func Init(c Config) {
	Local = New(Param{
		Addr:     c.Addr,
		Password: c.Password,
		MaxIdle:  c.MaxIdle,
		Db:       c.Db,
	})
}

type Param struct {
	Addr     string
	Password string
	MaxIdle  int
	Db       int
}

func New(param Param) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     param.MaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				param.Addr,
				redis.DialPassword(param.Password),
				redis.DialDatabase(param.Db),
			)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			if err != nil {
				log.Debugf("EROR:ping redis fail:%v", err)
			}
			return err
		},
	}
}
