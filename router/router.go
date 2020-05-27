package router

import (
	"github.com/coffemanfp/shachat/handlers"
	"github.com/gin-gonic/gin"
)

// NewRouter - Router.
func NewRouter() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/", handlers.HelloWorld)
	r.GET("/ping", handlers.PingPong)

	// Users
	r.GET("/users", handlers.GetUsers)

	return r
}
