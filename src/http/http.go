package http

import (
	"appletMessagesServer/src/g"

	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type (
	Dto struct {
		Msg  string      `json:"msg"`
		Ts   string      `json:"ts"` // 时间戳
		Data interface{} `json:"data"`
	}
)

func Start() {
	configApiRoutes()
	configW3Routes()
	// 静态资源请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(filepath.Join(g.Root, "/public"))).ServeHTTP(w, r)
	})

	Adress := g.ConfigTo().HTTP.Listen
	// start http server
	s := &http.Server{
		Addr:           Adress,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("http.Start ok, listening on", Adress)
	log.Fatalln(s.ListenAndServe())
}

func Stop() {
	log.Println("http.Stop ok")
}

func RenderImage200(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "image/jpg; charset=UTF-8")
	w.WriteHeader(200)
	w.Write(b)
}

func RenderText200(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.WriteHeader(200)
	w.Write([]byte(s))
}

func RenderText(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/text; charset=UTF-8")
	w.Write([]byte(s))
}

func RenderXml(w http.ResponseWriter, v interface{}) {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
	w.Write(bs)
}

func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	RenderJson(w, Dto{Msg: "success", Ts: time.Now().Format("20060102150405"), Data: data})
}

func RenderMsgJson(w http.ResponseWriter, msg string) {
	RenderJson(w, map[string]string{"msg": msg})
}

func AutoRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		RenderMsgJson(w, err.Error())
		return
	}
	RenderDataJson(w, data)
}

func StdRender(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		w.WriteHeader(400)
		RenderMsgJson(w, err.Error())
		return
	}
	RenderJson(w, data)
}
