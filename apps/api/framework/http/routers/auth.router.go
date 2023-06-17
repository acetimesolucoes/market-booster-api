package routers

import (
	"marketbooster/platform/authenticator"
	"marketbooster/platform/authenticator/callback"
	"marketbooster/platform/authenticator/login"
	"marketbooster/platform/authenticator/logout"
	"marketbooster/platform/authenticator/user"

	"github.com/gin-gonic/gin"
)

type AuthenticationRouter struct{}

func (auth *AuthenticationRouter) CreateRouter(rg *gin.RouterGroup) {

	a := new(authenticator.Authenticator)

	r := rg.Group("oauth")
	{
		r.GET("/login", login.Handler(a))
		r.GET("/callback", callback.Handler(a))
		r.GET("/user", user.Handler)
		r.GET("/logout", logout.Handler)
	}
}
