package model

import (
	"appletMessagesServer/src/g"
	"appletMessagesServer/src/util"

	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/imroc/log"
)

//AppletConfigValid 验证小程序配置参数
func AppletConfigValid(c *g.AppletConfig) {
	if c == nil {
		panic("[Warn] applet config not find")
	}
}

//AppletQueryParamsValid 小程序回调参数合法性检验
func AppletQueryParamsValid(m url.Values) {
	nonce := m.Get("nonce")
	timestamp := m.Get("timestamp")
	signature := m.Get("signature")
	msg_signature := m.Get("msg_signature")

	if nonce == "" {
		panic("nonce is nil")
	}
	if timestamp == "" {
		panic("timestamp is nil")
	}
	if signature == "" && msg_signature == "" {
		panic("signature and msg_signature is nil")
	}
}

//AppletSignValid 验证消息是否来自微信服务器
func AppletSignValid(c *g.AppletConfig, m url.Values) {
	nonce := m.Get("nonce")
	timestamp := m.Get("timestamp")
	signature := m.Get("signature")
	if util.Sign(c.Token, timestamp, nonce) == signature {
		return
	} else {
		log.Debug("signature not match")
		panic("signature not match")
	}
}

//AppletStrValid 验证加密类型是否一致
func AppletStrValid(v, w, e string) {
	if v != w {
		panic(e)
	}
}

//AppletMessageXMLValid xml解析
func AppletMessageXMLValid(req *http.Request, aesBody *AesRequestBody) {
	if err := xml.NewDecoder(req.Body).Decode(aesBody); err != nil {
		log.Info("[Warn] xml body", err)
		panic("xml body parse err")
	}
}

//AppletSignEncryptValid 加密模式的指纹验证方法
func AppletSignEncryptValid(c *g.AppletConfig, m url.Values, body string) {
	nonce := m.Get("nonce")
	timestamp := m.Get("timestamp")
	signature := m.Get("msg_signature")
	if util.MsgSign(c.Token, timestamp, nonce, body) == signature {
		return
	} else {
		panic("signature not match")
	}
}

//ProcessAppletText 处理小程序收到的文本消息
func ProcessAppletText(c *g.AppletConfig, mixedMsg *MixedMessage) string {
	//text := GetText(mixedMsg)
	text := mixedMsg.Content
	log.Info("text.Cotent:", text)
	if text == "天王盖地虎" {
		SendMessageText(c.AppId, mixedMsg.FromUserName, "宝塔镇河妖")
	} else {
		SendMessageText(c.AppId, mixedMsg.FromUserName, "test")
	}

	return ""
}

//ProcessAppletEvent 处理小程序收到的事件消息
func ProcessAppletEvent(c *g.AppletConfig, mixedMsg *MixedMessage) {
	switch mixedMsg.Event {
	case EventEnterTempsession:
		media_id := "xxxxxxx" //media_id是通过把媒体文件上传到微信服务器获得的
		SendMessageImage(c.AppId, mixedMsg.FromUserName, media_id)
		SendMessageText(c.AppId, mixedMsg.FromUserName, "您好，请长按二维码加入群聊")
	default:
		SendMessageText(c.AppId, mixedMsg.FromUserName, "您好，很高兴为您服务。")
	}
}
