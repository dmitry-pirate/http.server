package middleware

import (
	"github.com/basketforcode/http.server/app/services/cache"
	"github.com/basketforcode/http.server/app/services/store"
	"github.com/gin-gonic/gin"
)

func InjectMiddleware(store *store.Store, cache *cache.Redis) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		c.Set("cache", cache)
		c.Next()
	}
}
