package route

import (
	example_controller "foo/controller/example"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(group *gin.RouterGroup) {
	// do stuff here
	example := group.Group("/example")
	ctl := example_controller.New()
	{
		example.GET("/get", ctl.Get)
		example.GET("/set", ctl.Put)
	}
}
