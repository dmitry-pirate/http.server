package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vpn_api/app/config"
	"vpn_api/app/models"
	"vpn_api/app/store"
)

func AuthMiddleware(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		st, err := store.New(conf)
		if err != nil {
			panic(err)
		}
		token := c.GetHeader("Authorization")
		var user models.Users
		if err := st.GetConnection().Get(&user, "select id from users where id = (select id from user_token where token = ?) limit 1", token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "User with credentials is missing"})
			return
		}
		c.Next()
	}
}
