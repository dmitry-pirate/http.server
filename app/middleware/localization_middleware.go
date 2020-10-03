package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	LangEN = "en"
	LangRU = "ru"
	LangUK = "uk"
)

func LocalizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("Accept-Language")
		if isInSupportedLocales(locale) {
			c.Set("locale", locale)
		} else {
			c.Set("locale", LangEN)
		}
		c.Next()
	}
}

func isInSupportedLocales(locale string) bool {
	var locales = [3]string{LangEN, LangRU, LangUK}
	for _, l := range locales {
		if locale == l {
			return true
		}
	}
	return false
}
