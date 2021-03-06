package routers

/*
 * @Script: routers.go
 * @Author: pangxiaobo
 * @Email: 10846295@qq.com
 * @Create At: 2018-11-27 18:19:27
 * @Last Modified By: pangxiaobo
 * @Last Modified At: 2018-12-12 14:25:18
 * @Description: This is description.
 */

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_init/controller"
	"github.com/go_init/middleware"
)

var indexCtl = new(controller.IndexController)
var testCtl = new(controller.TestController)
var wsCtl = new(controller.WsController)
var mqCtl = new(controller.MqController)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	//router.Use(gin.Logger())

	router.GET("/", indexCtl.Welcome)
	router.NoRoute(indexCtl.Handle404)
	router.GET("/redis", testCtl.RedisTest) //redis 测试

	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.unclepang.com/")
	})
	router.POST("/exchange", func(c *gin.Context) {
		mqCtl.ExchangeHandler(c.Writer, c.Request)
	})
	router.POST("/queue/bind", func(c *gin.Context) {
		mqCtl.QueueBindHandler(c.Writer, c.Request)
	})
	router.GET("/queue", func(c *gin.Context) {
		mqCtl.QueueHandler(c.Writer, c.Request)
	}) //consume queue
	router.POST("/queue", func(c *gin.Context) {
		mqCtl.QueueHandler(c.Writer, c.Request)
	}) //declare queue
	router.DELETE("/queue", func(c *gin.Context) {
		mqCtl.QueueHandler(c.Writer, c.Request)
	}) //delete queue
	router.POST("/publish", func(c *gin.Context) {
		mqCtl.PublishHandler(c.Writer, c.Request)
	})
	router.GET("/ws", func(c *gin.Context) {
		wsCtl.WsHandler(c.Writer, c.Request)
	})

	v1 := router.Group("/v1")
	v1.Use(middleware.CORS(middleware.CORSOptions{}))
	{
		v1.GET("/test", testCtl.GetNick)
	}

	v2 := router.Group("/v2")
	v2.Use(middleware.CORS(middleware.CORSOptions{}))
	{
		v2.GET("/user", testCtl.GetUser)
		v2.POST("/user", testCtl.AddUser)
		v2.DELETE("/user", testCtl.DelUser)
		v2.PATCH("/user", testCtl.UptUser)
	}

	return router
}
