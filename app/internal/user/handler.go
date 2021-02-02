package user

import (
	"fmt"
	"github.com/basketforcode/http.server/app/internal/usertoken"
	"github.com/basketforcode/http.server/app/services/cache"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/basketforcode/http.server/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const infoCacheKey = "user_info:%s"

type handler struct {
	config *config.Config
	store  *store.Store
	cache  *cache.Redis
}

//New user handler struct
func NewHandler(store *store.Store, cache *cache.Redis, config *config.Config) *handler {
	return &handler{
		config: config,
		store:  store,
		cache:  cache,
	}
}

//Handle http user info
func (h *handler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Value("auth").(usertoken.UserToken)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "credentials does not match"})
			return
		}

		cacheKey := fmt.Sprintf(infoCacheKey, token.Token)
		usr, err := h.cache.Get(c, cacheKey)
		if err != nil {
			usr, err = NewRepo(h.store).FormatInfo(token)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": err.Error()})
				return
			}

			if err := h.cache.Set(c, cacheKey, usr, time.Hour*12); err != nil {
				log.Print(err)
			}
		}

		c.AsciiJSON(http.StatusOK, gin.H{
			"success": true,
			"data":    usr,
		})
	}
}
