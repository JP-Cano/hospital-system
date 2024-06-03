package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleError(c *gin.Context, err error, statusCode int) {
	if err != nil {
		log.Println(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
}
