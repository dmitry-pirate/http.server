package handlers

import (
	"github.com/dmitry-pirate/http.server/app/services/cache"
	"github.com/dmitry-pirate/http.server/app/services/store"
	"github.com/dmitry-pirate/http.server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func (handler *userHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"type": "success",
			"data": "",
		})
	}
}
