package handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var authBearerRegExp *regexp.Regexp = regexp.MustCompile("[B|b]earer (.*)")

func ValidateAPIKey(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := authBearerRegExp.FindStringSubmatch(c.GetHeader("Authorization"))
		logrus.Debug("RegExp Debug: ", result)
		if len(result) != 2 || result[1] != apiKey {
			errMessage := "APIKEY is not found or is incorrected"
			logrus.Warn(errMessage, ", Authorization Header: ", c.GetHeader("Authorization"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
			return
		}
		c.AddParam("valid_test", "true")
	}
}
