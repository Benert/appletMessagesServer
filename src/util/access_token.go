package util

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/imroc/log"
	"github.com/toolkits/net/httplib"
)

type AccessTokenInfo struct {
	ErrCode   int64  `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"` // 有效时间, seconds
}

//GetToken 去获取微信服务器获取access_token
func GetToken(appId, appSecret string) *AccessTokenInfo {

	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + url.QueryEscape(appId) +
		"&secret=" + url.QueryEscape(appSecret)

	r := httplib.Get(url).SetTimeout(3*time.Second, 1*time.Minute)
	resp, err := r.String()
	FailOnError(err, "EROR: refresh token")

	var token AccessTokenInfo

	if err = json.Unmarshal([]byte(resp), &token); err != nil {
		log.Debug("[ERROR] json ", err, resp)
		return nil
	}

	if token.ErrCode != 0 {
		return nil
	}

	return &token
}
