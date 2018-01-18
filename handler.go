package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokenMessage struct {
	Token string `json:"token" binding:"required"`
}

// HandlerFunc for /fcm/register. Expects a JSON body containing the fields of
// tokenMessage and adds any new tokens to the global map of Firebase Cloud
// Messaging tokens.
func registerNewFCMID(c *gin.Context) {
	var tm tokenMessage
	if err := c.BindJSON(&tm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := fcmTokenMap[tm.Token]; !ok {
		fcmTokenMap[tm.Token] = true
	}
}
