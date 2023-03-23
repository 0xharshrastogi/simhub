package route

import (
	"github.com/gin-gonic/gin"
	"github.com/harshrastogi/eagle/cmd/server/controller"
)

func New() *gin.Engine {
	router := gin.New()
	r := router.Group("/resource")
	{
		r.GET("/", controller.GetAllResources)
	}

	return router
}
