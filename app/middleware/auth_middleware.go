package middleware

import (
	"github.com/basketforcode/http.server/app/repositories"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		s, _ := c.Value("store").(*store.Store)

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Unauthorized"})
		}

		userRepo := repositories.NewUserTokenRepo(s)
		userToken, err := userRepo.Get(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Unauthorized"})
		}
		c.Set("auth", userToken)

		c.Next()
	}
}
