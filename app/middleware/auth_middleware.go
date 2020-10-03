package middleware

import (
	"github.com/basketforcode/http.server/app/repositories"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		userRepo := repositories.NewUserTokenRepo(store)
		userToken, err := userRepo.Get(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Credentials does not match"})
		}
		c.Set("auth", userToken)
		c.Next()
	}
}
