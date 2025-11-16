package router

import (
	"sc_police_api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//解决跨域
	r.Use(cors.Default())

	auth := r.Group("/api/auth")
	{
		auth.GET("checkToken", controllers.CheckToken)
	}

	user := r.Group("/api/user")
	{
		user.POST("createUser", controllers.CreateUser)

	}

	return r
}
