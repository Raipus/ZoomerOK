package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(release bool) *gin.Engine {
	if release {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.TestMode)
	}
	router := gin.Default()
	router.Use(gin.Logger())
	return router
}
