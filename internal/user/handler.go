package user

import (
	"fmt"
	"github.com/basketforcode/http.server/internal/usertoken"
	"github.com/basketforcode/http.server/pkg/cache"
	"github.com/basketforcode/http.server/pkg/config"
	"github.com/basketforcode/http.server/pkg/store"
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

func NewHandler(store *store.Store, cache *cache.Redis, config *config.Config) *handler {
	return &handler{
		config: config,
		store:  store,
		cache:  cache,
	}
}

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
			dbc := h.store.SlaveConnection()

			usr, err = NewRepo(dbc).FormatInfo(token)

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
