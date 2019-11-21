package routers

import (
	UserController "GolangLearning/controller/user"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	user := new(UserController.User)

	v1 := engine.Group("/")
	{
		v1.GET("/user/queryById/:id", user.QueryById)
		v1.GET("/user/delete", user.Delete)
		v1.GET("/user/queryUser", user.QueryUser)
		v1.POST("/user/create", user.Create)
		v1.POST("/user/update", user.Update)
	}

	return engine
}
