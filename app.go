package main

import (
	"myapp/wxmp"
	"net/http"
)

// 文本消息的处理函数
func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content          // 获取用户发送的消息
	w.ReplyText(txt)          // 回复一条文本消息
	w.PostText("Post:" + txt) // 发送一条文本消息
}

// 关注事件的处理函数
func Subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}

func main() {
	// my-token 验证微信公众平台的Token
	// app-id, app-secret用于高级API调用。
	// 如果仅使用接收/回复消息，则可以不填写，使用下面语句
	// mux := weixin.New("my-token", "", "")
	mux := weixin.New("lzbgt1", "wxb7d39933d3a607af", "c115d488b48d0b27c8a9b10605177cf7")
	// 注册文本消息的处理函数
	mux.HandleFunc(weixin.MsgTypeText, Echo)
	// 注册关注事件的处理函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, Subscribe)
	http.Handle("/weixin/access", mux) // 注册接收微信服务器数据的接口URI
	http.ListenAndServe(":80", nil)    // 启动接收微信数据服务器
}
