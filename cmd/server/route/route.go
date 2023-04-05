package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/harshrastogi/eagle/cmd/server/controller"
)

func New() *gin.Engine {
	router := gin.New()
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true

	router.Use(cors.New(c))

	r := router.Group("/resource")
	{
		r.GET("/", controller.GetAllResources)
	}

	return router
}
