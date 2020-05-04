package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vpn_api/app/config"
	"vpn_api/app/models"
	"vpn_api/app/requests"
	"vpn_api/app/store"
)

//HandlePing ...
func HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(requests.PacRequest)
		if err := c.Bind(request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Request is not parsed"})
			return
		}
		conf := config.NewConfig()
		s, err := store.New(conf)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Config is not parsed"})
			return
		}
		defer func() {
			if err := s.Close(); err != nil {
				panic(err)
			}
		}()

		var proxies []models.Proxies
		if err := s.GetConnection().Select(&proxies, "select * from proxies where premium = ? and user_version = ?", 1, "UVPNv2"); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Server is not found"})
			return
		}
		bypass := models.GetAllBypass(s)
		proxiesJson := models.PrepareProxiesPingResults(s, &proxies)
		c.JSON(http.StatusOK, gin.H{"mode": "vpn", "currentProxyServer": proxiesJson, "bypass": bypass})
	}
}
