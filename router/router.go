package router

import (
	"MytheresaChallenge/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter() (router *gin.Engine) {
	router = gin.Default()

	//generate prefix /api
	api := router.Group("api")
	{
		api.GET("products", controllers.GetProducts)
	}

	return
}