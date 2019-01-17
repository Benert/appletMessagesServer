package http

import (
	"appletMessagesServer/src/g"
	"appletMessagesServer/src/model"
	"encoding/xml"
	"io/ioutil"

	"net/http"
	"net/url"

	"github.com/imroc/log"
)

var (
	appId string
)

//configApiRoutes api 接口 返回josn、xml
func configApiRoutes() {
	http.HandleFunc("/play/debug", func(w http.ResponseWriter, req *http.Request) {
		log.Info("request /play/debug")
		StdRender(w, g.VERSION, nil)
	})

	AppletArr := g.ConfigTo().Applet
	for _, v := range AppletArr {
		http.HandleFunc("/"+v.ActionAddres, func(w http.ResponseWriter, r *http.Request) {
			// 捕获异常
			defer func() {
				if r := recover(); r != nil {
					log.Debugf("Runtime error caught: %v", r)
					w.WriteHeader(400)
					w.Write([]byte(""))
					return
				}
			}()
			log.Info("request url：", v.ActionAddres)

			var queryValues url.Values

			queryValues, _ = url.ParseQuery(r.URL.RawQuery)
			switch r.Method {
			case "GET":
				{
					model.AppletSignValid(v, queryValues)
					RenderText200(w, queryValues.Get("echostr"))
					return
				}
			case "POST":
				{
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Debug("read body err:,", err)
						return
					}
					log.Infof("report body:%s", body)
					var commonBody model.MixedMessage
					if err := xml.Unmarshal(body, &commonBody); err != nil {
						log.Debug("[Warn] body Unmarshal", err)
						w.WriteHeader(400)
						return
					}

					switch commonBody.MsgType {
					// text
					case model.MsgTypeText:
						{
							model.ProcessAppletText(v, &commonBody) // 文本消息的处理逻辑
						}

					// event
					case model.MsgTypeEvent:
						{
							model.ProcessAppletEvent(v, &commonBody)
						}

					}
					w.WriteHeader(200)
					RenderText(w, "")
				}
			}

		})
	}
}
