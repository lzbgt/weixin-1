package main

import (
	"myapp/weixin"
	"net/http"
    "log"
)


func deleteMenu(wx *weixin.Weixin) {
    err := wx.DeleteMenu()
    if err != nil {
        log.Print(err)
    }
}

var menu_keys = []string{
"MENU_KEY_CV_NEW",
"MENU_KEY_CV_CALIX",
"MENU_KEY_CV_MOTO",
"MENU_KEY_CV_ZTE",
}

var menu_cvs = []string{
    "期待与能与您合作... TEL:17092500181",
    "CALIX:\n参与CALIX EXA COMPASS项目的开发维护.\n" +"其核心是将传统的硬件管理软件向云端WEB迁移从而涉及一系列的开发和改造\n1)后台改造:开发系统API接口, 开发数据同步系统, 大规模并发改造\n2)开发WEB前端\n本人有幸参与了整个项目,期间接触了一系列的平台,工具和语言以及系统架构包括:SOAP, WEBSOCKET, REST, GO, C/C++, JAVA, SPRING, JBOSS, PERL, RUBY, NODE.JS, ZMQ",
    "MOTOROLA/NSN:\n主要负责WiMAX/LTE通信系统中BTS/AP的高可靠性软件开发: C/C++",
    "ZTE:\n主要负责BTS边界网关软件开发: C/C++",
}

func createMenu(wx *weixin.Weixin){
    deleteMenu(wx)
    menu := &weixin.Menu{make([]weixin.MenuButton, 2)}
    menu.Buttons[0].Name = "我的简历"
    menu.Buttons[0].SubButtons = make([]weixin.MenuButton, 4)
    menu.Buttons[0].SubButtons[0].Name = "您的公司?"
    menu.Buttons[0].SubButtons[0].Type = weixin.MenuButtonTypeKey
    menu.Buttons[0].SubButtons[0].Key = menu_keys[0]
    menu.Buttons[0].SubButtons[1].Name = "CALIX"
    menu.Buttons[0].SubButtons[1].Type = weixin.MenuButtonTypeKey
    menu.Buttons[0].SubButtons[1].Key = menu_keys[1]
    menu.Buttons[0].SubButtons[2].Name = "MOTOROLA/NSN"
    menu.Buttons[0].SubButtons[2].Type = weixin.MenuButtonTypeKey
    menu.Buttons[0].SubButtons[2].Key = menu_keys[2]
    menu.Buttons[0].SubButtons[3].Name = "ZTE"
    menu.Buttons[0].SubButtons[3].Type = weixin.MenuButtonTypeKey
    menu.Buttons[0].SubButtons[3].Key = menu_keys[3]
    menu.Buttons[1].Name = "关于"
    menu.Buttons[1].SubButtons = make([]weixin.MenuButton, 2)
    menu.Buttons[1].SubButtons[0].Name = "更多"
    menu.Buttons[1].SubButtons[0].Type = weixin.MenuButtonTypeKey
    menu.Buttons[1].SubButtons[0].Key = "MENU_KEY_MORE"
    menu.Buttons[1].SubButtons[1].Name = "联系方式"
    menu.Buttons[1].SubButtons[1].Type = weixin.MenuButtonTypeKey
    menu.Buttons[1].SubButtons[1].Key = "MENU_KEY_CONTACT"
    err := wx.CreateMenu(menu)
    if err != nil {
        log.Print(err)
    }
}

func MenuHandler(cmd string, wx *weixin.Weixin) {
    if cmd == "cm" {
        createMenu(wx)
    } else if cmd == "dm" {
        deleteMenu(wx)
    }
}

func Echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content          // 获取用户发送的消息
	w.ReplyText(txt)          // 回复一条文本消息
	//w.PostText("Post:" + txt) // 发送一条文本消息
    MenuHandler(r.Content, w.GetWeixin())   
}


var  artis = []weixin.Article{
    {"我的近照","","http://103.27.109.177/images/thumb_head.jpg",""},
}

func ClickHandler(w weixin.ResponseWriter, r *weixin.Request) {
    switch r.EventKey {
    case "MENU_KEY_RESUME":
        w.ReplyNews(artis)
    case "MENU_KEY_CONTACT":
        w.ReplyText("联系人: 陆宗宝, 17092500181, weixin: 41552136\n本后台使用GO语言开发")
    default:
        for i, t := range menu_keys {
            if t == r.EventKey {
                w.ReplyText(menu_cvs[i]);
                return
            }
        }
        w.ReplyText("你点击了菜单:\n\t" + r.EventKey + "\n此功能正在建设中...")
    }
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
    mux.HandleFunc(weixin.MsgTypeEventClick, ClickHandler)
	http.Handle("/weixin/access", mux) // 注册接收微信服务器数据的接口URI
	http.ListenAndServe(":80", nil)    // 启动接收微信数据服务器
}
