package router

import (
	"testing/src/controller"
	docs "testing/src/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/vi"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("user")
		{
			//增加用户User
			userGroup.POST("/users", controller.CreateUser)
			//查看所有的User
			userGroup.GET("/users", controller.GetUserList)
			//修改某个User
			userGroup.PUT("/users/:id", controller.UpdateUser)
			// //删除某个User
			// userGroup.DELETE("/users/:id", controller.DeleteUserById)
		}

		// testGroup := v1.Group("test")
		// {
		// 	//增加用户User
		// 	testGroup.POST("", controller.CreateModel)
		// 	//查看所有的User
		// 	testGroup.GET("", controller.GetModel)
		// 	//修改某个User
		// 	testGroup.PUT(":id", controller.UpdateModel)
		// 	// //删除某个User
		// 	// userGroup.DELETE("/users/:id", controller.DeleteUserById)
		// }
	}

	return r
}
