package routers

import (
	"marketbooster/framework/http/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (e *UserRouter) CreateRouter(rg *gin.RouterGroup) {

	r := rg.Group("users")
	{
		controller := new(controllers.UserController)

		r.GET("/", controller.FindAll)
		r.GET("/:id", controller.FindById)
		r.POST("/", controller.Create)
		r.PUT("/:id", controller.Update)
		r.DELETE("/:id", controller.Delete)
	}
}
