package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ValidateAPIKey(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != apiKey {
			errMessage := "APIKEY not found"
			logrus.Warn(errMessage, ", APIKEY: ", c.GetHeader("Authorization"))
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
			return
		}
	}
}
