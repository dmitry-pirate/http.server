package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlePing ...
func HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"mode": "vpn"})
	}
}
