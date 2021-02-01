package main

import (
	"fmt"
	"gin-demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/pkg/log"
)

func main ()() {
	// conf Init
	if err := conf.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}
	// log Init
	log.Init(conf.Conf.Log)
	defer log.Close()

	log.Info("-------------project start-------------")
	ginTest := gin.Default()
	ginEngine := GinRouter(ginTest)

	ginTest.GET("/ping", test)

	ginEngine.Run(":8080")
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