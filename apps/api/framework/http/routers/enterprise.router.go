package enterprise_router

// import (
// 	"acetime.com.br/business-crm/apps/api/application/controllers"
// 	"github.com/gin-gonic/gin"
// )

// type EnterpriseRouter struct{}

// func (e EnterpriseRouter) CreateRouter(routerGroup *gin.RouterGroup) gin.RouterGroup {
// 	routerGroup.Group("/v1")
// 	{
// 		enterprises := v1.Group("enterprises")
// 		{
// 			controller := new(controllers.EnterpriseController)
// 			// Enterprises
// 			enterprises.GET("/", controller.FindAll)
// 			enterprises.GET("/:id", controller.FindById)
// 			enterprises.POST("/", controller.Create)
// 			enterprises.PUT("/:id", controller.Update)
// 			enterprises.DELETE("/:id", controller.Delete)
// 		}
// 	}
// }
