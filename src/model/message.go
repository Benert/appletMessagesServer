package model

const (
	// 小程序服务器推送过来的消息类型
	MsgTypeEvent  = "event"           //进入会话事件
	MsgTypeText   = "text"            // 文本消息
	MsgTypeImage  = "image"           // 图片消息
	MsgTypeApplet = "miniprogrampage" //小程序卡片信息
	MsgTypeLink   = "link"            // 图文消息
)

// 安全模式, 微信服务器推送过来的 http body
type AesRequestBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ToUserName   string `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string `xml:"Encrypt"    json:"Encrypt"`
}

// 一般模式下面，微信服务器推送过来的 http body
type NormalRequestBody struct {
	XMLName    struct{} `xml : "xml" json:"-"`
	ToUserName string   `xml:"ToUserName" json:"ToUserName"`
}

// 安全模式下回复消息的 http body
type AesResponseBody struct {
	XMLName struct{} `xml:"xml" json:"-"`

	EncryptedMsg string `xml:"Encrypt"      json:"Encrypt"`
	MsgSignature string `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64  `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string `xml:"Nonce"        json:"Nonce"`
}

// 微信服务器推送过来的消息(事件)通用的消息头
type MessageHeader struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
	MsgId        int64  `xml:"MsgId"        json:"MsgId"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MessageHeader

	MsgId        int64  `xml:"MsgId"        json:"MsgId"`
	Content      string `xml:"Content"      json:"Content"`
	PicUrl       string `xml:"PicUrl"       json:"PicUrl"`
	MediaId      string `xml:"MediaId"      json:"MediaId"`
	Title        string `xml:"Title"        json:"Title"`
	AppId        string `xml:"AppId"        json:"AppId"`
	PagePath     string `xml:"PagePath"     json:"PagePath"`
	ThumbUrl     string `xml:"ThumbUrl"     json:"ThumbUrl"`
	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"`
	Event        string `xml:"Event"        json:"Event"`
	SessionFrom  string `xml:"SessionFrom"  json:"SessionFrom"`
}

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MessageHeader

	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func GetText(msg *MixedMessage) *Text {
	return &Text{
		MessageHeader: msg.MessageHeader,
		Content:       msg.Content,
	}
}

//Image 图片消息结构
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MessageHeader
	MsgId   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	MediaId string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	PicURL  string `xml:"PicUrl"  json:"PicUrl"`  // 图片链接
}

//GetImage 获取图片消息
func GetImage(msg *MixedMessage) *Image {
	return &Image{
		MessageHeader: msg.MessageHeader,
		MediaId:       msg.MediaId,
		PicURL:        msg.PicUrl,
	}
}
