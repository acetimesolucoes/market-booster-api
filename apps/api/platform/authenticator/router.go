package authenticator

import (
	"encoding/gob"
	"marketbooster/platform/authenticator/callback"
	"marketbooster/platform/authenticator/login"
	"marketbooster/platform/authenticator/logout"
	"marketbooster/platform/authenticator/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

// New registers the routes and returns the router.
func (r *AuthRouter) New(auth *Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
