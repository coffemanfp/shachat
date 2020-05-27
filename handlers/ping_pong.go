package handlers

import "github.com/gin-gonic/gin"

// PingPong - Sends a pong to the client.
func PingPong(c *gin.Context) {
	c.String(200, "pong")
}
