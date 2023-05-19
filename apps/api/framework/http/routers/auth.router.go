package routers

import (
	"marketbooster/framework/http/controllers"
	"github.com/gin-gonic/gin"
)

type AuthenticationRouter struct{}

func (auth *AuthenticationRouter) CreateRouter(rg *gin.RouterGroup) {

	r := rg.Group("oauth")
	{
		controller := new(controllers.EnterpriseController)

		r.GET("/login", login.Handler(auth))
		r.GET("/callback", callback.Handler(auth))
		r.GET("/user", user.Handler)
		r.GET("/logout", logout.Handler)
	}
}
