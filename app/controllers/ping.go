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
			panic(err)
		}
		conf := config.NewConfig()
		s, err := store.New(conf)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := s.Close(); err != nil {
				panic(err)
			}
		}()

		var proxies []models.Proxies
		if err := s.GetConnection().Select(&proxies, "select * from proxies where premium = ?", 1); err != nil {
			panic(err)
		}
		proxiesJson := models.PrepareProxiesPingResults(s, &proxies)
		c.JSON(http.StatusOK, gin.H{"mode": "vpn", "currentProxyServer": proxiesJson})
	}
}
