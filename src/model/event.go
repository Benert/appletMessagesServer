package model

const (
	// 微信服务器推送过来的事件类型
	EventEnterTempsession = "user_enter_tempsession" //进入会话事件
)

//进入会话事件
type EnterTempsessionEvent struct {
	XMLName struct{} `xml:"xml"    json:"-"`
	MessageHeader
	Event string `xml:"Event"      json:"Event"`
}

func GetSessionFromEvent(msg *MixedMessage) *EnterTempsessionEvent {
	return &EnterTempsessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
	}
}
