package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)
import "github.com/gin-gonic/gin"

var (
	defaultUpgrader = websocket.Upgrader{}
)

//websocket server
func main() {
	engine := gin.Default()
	engine.Use(gin.Recovery(), gin.Logger())
	g := engine.Group("/test", func(context *gin.Context) {
		fmt.Printf("group use middleware.......\n")
	})
	g.Use(func(context *gin.Context) {
		fmt.Printf("sub-path in group use middleware.......\n")
	})
	g.GET("/websocket", WebsocketTest)
	fmt.Printf("listen :7000...\n")
	engine.Run(":7000")
}
func WebsocketTest(ctx *gin.Context) {
	c, err := defaultUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("websocket ugrade failed...\n")
		return
	}
	i := 1
	defer func() {
		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}()

	for {
		c.SetReadDeadline(time.Now().Add(time.Second * 3))
		_, p, err := c.ReadMessage()
		if err != nil || string(p) == "end" {
			return
		}
		rs := rsp{
			Src:     string(p),
			Dest:    strconv.Itoa(i) + ":" + string(p),
			ErrCode: 0,
			ErrMsg:  "ok",
		}
		c.SetWriteDeadline(time.Now().Add(time.Second * 3))
		c.WriteJSON(&rs)
		i++
	}
}

type rsp struct {
	Src     string
	Dest    string
	ErrCode int
	ErrMsg  string
}
