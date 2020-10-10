package handlers

import (
	"github.com/basketforcode/http.server/app/models"
	"github.com/basketforcode/http.server/app/repositories"
	"github.com/basketforcode/http.server/app/services/cache"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/basketforcode/http.server/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const infoCacheKey = "user_info:"

type userHandler struct {
	config *config.Config
	store  *store.Store
	cache  *cache.Redis
}

//user userHandler struct
func NewUserHandler(store *store.Store, cache *cache.Redis, config *config.Config) *userHandler {
	return &userHandler{
		config: config,
		store:  store,
		cache:  cache,
	}
}

//HandleInfo is http handler
func (h *userHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Value("auth").(models.UserToken)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "Credentials does not match"})
			return
		}

		cacheKey := infoCacheKey + token.Token
		usr, err := h.cache.Get(c, cacheKey)
		if err != nil {
			r := repositories.NewUserRepo(h.store)
			usr, err = r.GetFormattedInfo(token)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": err.Error()})
				return
			}

			_ = h.cache.Set(c, cacheKey, usr, time.Hour*12)
		}

		c.AsciiJSON(http.StatusOK, gin.H{
			"success": true,
			"data":    usr,
		})
	}
}
