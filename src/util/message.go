package util

const (
	MsgTypeText   = "text"  // 文本消息
	MsgTypeImage  = "image" // 图片消息
	MsgTypeNews   = "link"  // 图文消息
	MsgTypeApplet = "miniprogrampage"

	ErrCodeOK                 = 0
	ErrMsgOk                  = "success"
	ErrCodeInvalidCredential  = 40001 // access_token 过期(无效)返回这个错误
	ErrCodeAccessTokenExpired = 42001 // access_token 过期(无效)返回这个错误(maybe!!!)
)

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int64  `json:"msgid"`
}

type MessageHeader struct {
	ToUser  string `json:"touser"` // 接收方 OpenID
	MsgType string `json:"msgtype"`
}

type AppletMessage struct {
	MessageHeader
	MiniPage MiniProgramPage `json:"miniprogrampage"`
}

//MiniProgramPage 小程序消息体
type MiniProgramPage struct {
	Title        string `json:"title,omitempty"`
	Appid        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

//Text 文本消息
type Text struct {
	MessageHeader

	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`
}

//NewText 新建文本消息.
func NewText(token, toUser, content string) (text *Text) {
	text = &Text{
		MessageHeader: MessageHeader{
			ToUser:  toUser,
			MsgType: MsgTypeText,
		},
	}
	text.Text.Content = content

	return
}

//Image 图片消息
type Image struct {
	MessageHeader

	Image struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"image"`
}

//NewImage 新建图片消息.
func NewImage(token, toUser, mediaId string) (image *Image) {
	image = &Image{
		MessageHeader: MessageHeader{
			ToUser:  toUser,
			MsgType: MsgTypeImage,
		},
	}
	image.Image.MediaId = mediaId

	return
}
