package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld - Greeting.
func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
	return
}
