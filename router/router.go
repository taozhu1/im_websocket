package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"im-websocket/api"
	docs "im-websocket/docs"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 健康检测
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	// 首页
	r.GET("index", api.GetIndex)
	// user
	user := r.Group("/user")
	{
		user.POST("/register", api.UserRegister)
		user.GET("/getUserList", api.GetUserList)
		user.DELETE("/deleteUser", api.DeleteUser)
		user.PUT("/updateUser", api.UpdateUser)
		user.POST("/login", api.UserLogin)
		user.GET("/msg", api.WsHandler)
	}

	return r
}
