package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

const (
	HEART_BEAT              = 5005
	RECV_TXT_MSG            = 1
	RECV_PIC_MSG            = 3
	USER_LIST               = 5000
	GET_USER_LIST_SUCCSESS  = 5001
	GET_USER_LIST_FAIL      = 5002
	TXT_MSG                 = 555
	PIC_MSG                 = 500
	AT_MSG                  = 550
	CHATROOM_MEMBER         = 5010
	CHATROOM_MEMBER_NICK    = 5020
	PERSONAL_INFO           = 6500
	DEBUG_SWITCH            = 6000
	PERSONAL_DETAIL         = 6550
	DESTROY_ALL             = 9999
	NEW_FRIEND_REQUEST      = 37    //微信好友请求消息
	AGREE_TO_FRIEND_REQUEST = 10000 //同意微信好友请求消息
	ATTATCH_FILE            = 5003
)

type Message struct {
	Id       int
	Mtype    uint `json:"type"`
	Wxid     string
	Roomid   string
	Content  string
	Nickname string
	Ext      string
}

func main() {
	// init
	// schema – can be ws:// or wss://
	// host, port – WebSocket server
	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:5555",
		Path:   "/",
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect err: %v", err)
	}

	// send message
	msg, err := json.Marshal(&Message{Id: 0, Mtype: TXT_MSG, Wxid: "", Roomid: "", Content: " a new socket has connected.", Nickname: "doogie", Ext: ""})
	if err != nil {
		panic(err)
	}
	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "send err: %v", err)
	}
	// receive message
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println(message)
	}
}
