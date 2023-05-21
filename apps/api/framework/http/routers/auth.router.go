package routers

import (
	"marketbooster/platform/authenticator/logout"
	"marketbooster/platform/authenticator/user"

	"github.com/gin-gonic/gin"
)

type AuthenticationRouter struct{}

func (auth *AuthenticationRouter) CreateRouter(rg *gin.RouterGroup) {

	r := rg.Group("oauth")
	{
		// r.GET("/login", login.Handler(auth))
		// r.GET("/callback", callback.Handler(auth))
		r.GET("/user", user.Handler)
		r.GET("/logout", logout.Handler)
	}
}
