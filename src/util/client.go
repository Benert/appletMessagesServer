package util

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/imroc/log"
	"github.com/toolkits/net/httplib"
)

func Send(msg interface{}, access_token string) (err error) {

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + url.QueryEscape(access_token)

	req := httplib.Post(incompleteURL).SetTimeout(3*time.Second, 1*time.Minute)
	req.Body(msg)
	resp, err := req.String()

	log.Info(msg, resp)

	if err != nil {
		log.Debug("[ERROR]", err)
		return err
	}

	var result Error
	err = json.Unmarshal([]byte(resp), &result)
	if result.ErrCode != ErrCodeOK {
		log.Debug("[ERROR]", result)
		return
	}
	return
}
