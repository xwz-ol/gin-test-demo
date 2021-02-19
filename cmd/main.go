package main

import (
	"fmt"
	"gin-test-demo/conf"
	"gin-test-demo/dao"
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

	// New Dao And init
	dao.New(conf.Conf)
	log.Info("-------------project start-------------")
	log.V(1).Info("xxxxxxxxxxxxxxxxxxxxxxxxxxxxtestxxxxxxxxxxx")
	ginTest := gin.Default()
	ginEngine := GinRouter(ginTest)

	ginTest.GET("/ping", test)

	ginEngine.Run(conf.Conf.Web.Addr)
}
// test .
func test(context *gin.Context) {
	log.V(1).Info("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
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