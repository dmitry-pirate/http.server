package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vpn_api/app/config"
	"vpn_api/app/models"
	"vpn_api/app/requests"
	"vpn_api/app/store"
)

//HandlePac ...
func HandlePac() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := new(requests.PacRequest)
		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err.Error()})
			return
		}
		conf := config.NewConfig()
		st, err := store.New(conf)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Config is not parsed"})
			return
		}
		defer func() {
			if err := st.Close(); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": err.Error()})
				return
			}
		}()
		var country models.Countries
		if err := st.GetConnection().Get(&country, "select * from countries where country_code = ? limit 1", request.Country); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Country code is not found"})
			return
		}
		var proxies models.Proxies
		if err := st.GetConnection().Get(&proxies, "select * from proxies where country_id = ? and premium = ? limit 1", country.Id, 1); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "Server is not found with country code"})
			return
		}
		res := models.PrepareProxiesPacResults(st, &proxies, &country)
		c.JSON(http.StatusOK, gin.H{"mode": "vpn", "servers": res})
	}
}
