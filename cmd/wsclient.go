package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func main() {
	defaultDialer := websocket.DefaultDialer
	c, _, err := defaultDialer.Dial("ws://127.0.0.1:7000/test/websocket", nil)
	if err != nil {
		fmt.Printf("dial ws error(%v)\n", err)
		return
	}
	testSilce := []string{
		"今天是",
		"2020年8月8日，",
		"天气多云，",
		"最高温度33度。",
		"end", //最后发送end
	}
	for _, v := range testSilce {
		c.WriteMessage(websocket.TextMessage, []byte(v))
		_, p, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
		fmt.Printf("%v\n", string(p))
	}
}
