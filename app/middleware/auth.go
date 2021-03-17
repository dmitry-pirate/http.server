package middleware

import (
	"github.com/basketforcode/http.server/app/internal/usertoken"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(store *store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		//TODO: ??? connection must by once in context but context is bad container for db connection because anyone can change that connection
		userToken, err := usertoken.NewRepo(store.SlaveConnection()).Get(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "credentials does not match"})
		}
		c.Set("auth", userToken)
		c.Next()
	}
}
