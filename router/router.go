package router

import (
	"rest-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)

	router.Run("localhost: 3005")
}
