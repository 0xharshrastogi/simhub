package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	resourcecontroller "github.com/harshrastogi/eagle/cmd/server/controller/resource_controller"
)

func New() *gin.Engine {
	router := gin.New()
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true

	router.Use(cors.New(c))

	r := router.Group("/:realm/resource")
	{
		r.GET("/", resourcecontroller.GetAllResources)
	}

	return router
}
