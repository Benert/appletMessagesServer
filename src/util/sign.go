package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

// 微信公众号 明文模式/URL认证 签名
func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信公众号/企业号 密文模式消息签名
func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(encryptedMsg))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// 微信 js-sdk wx.config 的参数签名.
func WXConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	n := len("jsapi_ticket=") + len(jsapiTicket) +
		len("&noncestr=") + len(nonceStr) +
		len("&timestamp=") + len(timestamp) +
		len("&url=") + len(url)

	buf := make([]byte, 0, n)

	buf = append(buf, "jsapi_ticket="...)
	buf = append(buf, jsapiTicket...)
	buf = append(buf, "&noncestr="...)
	buf = append(buf, nonceStr...)
	buf = append(buf, "&timestamp="...)
	buf = append(buf, timestamp...)
	buf = append(buf, "&url="...)
	buf = append(buf, url...)

	//log.Println(string(buf))
	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}
func OpenSign(token, timestamp, nonce, appid string) (signature string) {
	// token + timestamp + nonce + appid
	strs := sort.StringSlice{token, timestamp, nonce, appid}
	//strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce)+len(appid))

	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)
	buf = append(buf, strs[3]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// MessageSign 发短信验证token
func MessageSign(str string) string {
	param := "b5f1399bda56a"
	h := md5.New()
	h.Write([]byte(str + param)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
