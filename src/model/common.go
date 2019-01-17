package model

import (
	"appletMessagesServer/src/util"
	"bytes"
	"encoding/json"
	"strings"
)

//  -------------   客服消息 小程序接口 ------------

//SendMessageText 给用户发送普通文本消息  客服消息接口
func SendMessageText(appid, openid, content string) {
	token := GetRedisAccessToken("applet_token_" + appid)
	obj := util.NewText(appid, openid, content)

	buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	buf.Reset()
	json.NewEncoder(buf).Encode(obj)
	tmpjson := buf.String()
	tmpjson = strings.Replace(tmpjson, "\\u0026", "<", -1)  //将json解析后的url中的"\u0026" 替换成"&"
	tmpjson = strings.Replace(tmpjson, "\\u003c", "<", -1)  //将json解析后的url中的"\u003c" 替换成"<"
	jsonNew := strings.Replace(tmpjson, "\\u003e", ">", -1) //将json解析后的url中的"\u003e" 替换成">"

	go util.Send(jsonNew, token)
}

//SendMessageImage 给用户发送图片消息 media_id上传到微信公众平台的素材
func SendMessageImage(appid, openid, media_id string) {
	token := GetRedisAccessToken("applet_token_" + appid)
	obj := util.NewImage(appid, openid, media_id)
	buf := bytes.NewBuffer(make([]byte, 0, 16<<10))
	buf.Reset()
	json.NewEncoder(buf).Encode(obj)
	tmpjson := buf.String()
	tmpjson = strings.Replace(tmpjson, "\\u0026", "<", -1)  //将json解析后的url中的"\u0026" 替换成"&"
	tmpjson = strings.Replace(tmpjson, "\\u003c", "<", -1)  //将json解析后的url中的"\u003c" 替换成"<"
	jsonNew := strings.Replace(tmpjson, "\\u003e", ">", -1) //将json解析后的url中的"\u003e" 替换成">"

	go util.Send(jsonNew, token)
}
