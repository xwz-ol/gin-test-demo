package main

import (
	"encoding/json"
	"fmt"
	"gin-test-demo/conf"
	"gin-test-demo/dao"
	"gin-test-demo/model"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/pkg/log"
	"golang.org/x/net/websocket"
	"net/http"
	"time"

)



// 全局信息
var datas model.Datas
var users map[*websocket.Conn]string

func main ()() {
	// conf Init
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// log Init
	log.Init(conf.Conf.Log)
	defer log.Close()
	dao.New(conf.Conf)
	log.Info("-------------project start-------------")
	ginTest := gin.Default()
	ginTest.LoadHTMLGlob("./html/*")
	ginEngine := GinRouter(ginTest)

	fmt.Println("启动时间: ", time.Now())

	// 初始化数据
	datas = model.Datas{}
	users = make(map[*websocket.Conn]string)

	// 渲染页面
	ginTest.GET("/", index)

	// 监听socket方法
		ginTest.GET("/webSocket", WebsocketStart(WebSocket))

		//// 监听8080端口
		//http.ListenAndServe(":8889", nil)

		ginTest.GET("/ping", test)

		ginEngine.Run(conf.Conf.Web.Addr)
	}
	// test .
	func test(context *gin.Context) {
		context.JSON(200,gin.H{
			"message": "Hello,world!",
	})
}

func GinRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/test")
	rr.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})

	rr = r.Group("/a")
	Routers(rr)

	return r
}

func Routers(r *gin.RouterGroup) {
	rr := r.Group("/b")
	rr.GET("/second", Function)

	return
}

func Function(c *gin.Context) {
	input :=  "first gin test demo"

	c.JSON(200,gin.H{
		"message": input,
	})
}

func index(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}

func WebsocketStart (ws websocket.Handler) gin.HandlerFunc{
	return func(c *gin.Context){
		if c.IsWebsocket() {
			ws.ServeHTTP(c.Writer,c.Request)
		}else {
			_, _ = c.Writer.WriteString("----NOT A Websocket----")
		}
	}
}
func WebSocket(ws *websocket.Conn)  {
	var message model.Message
	var data string
	for {
		// 接收数据
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			// 移除出错的连接
			delete(users, ws)
			fmt.Println("连接异常")
			break
		}
		// 解析信息
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			fmt.Println("解析数据异常")
		}

		// 添加新用户到map中,已经存在的用户不必添加
		if _, ok := users[ws]; !ok {
			users[ws] = message.Username
			// 添加用户到全局信息
			datas.Users = append(datas.Users, model.UserInfo{Username:message.Username})
		}
		// 添加聊天记录到全局信息
		datas.Messages = append(datas.Messages, message)


		// 通过webSocket将当前信息分发
		for key := range users{
			err := websocket.Message.Send(key, data)
			if err != nil{
				// 移除出错的连接
				delete(users, key)
				fmt.Println("发送出错: " + err.Error())
				break
			}
		}
	}
}