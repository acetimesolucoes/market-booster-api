package routers

import (
	"github.com/acetimesolutions/marketbooster/framework/http/controllers"
	"github.com/gin-gonic/gin"
)

type EnterpriseRouter struct{}

func (e *EnterpriseRouter) CreateRouter(rg *gin.RouterGroup) {

	r := rg.Group("enterprises")
	{
		controller := new(controllers.EnterpriseController)

		r.GET("/", controller.FindAll)
		r.GET("/:id", controller.FindById)
		r.POST("/", controller.Create)
		r.PUT("/:id", controller.Update)
		r.DELETE("/:id", controller.Delete)
	}
}
