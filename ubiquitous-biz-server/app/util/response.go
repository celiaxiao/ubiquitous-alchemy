package util

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorJSON(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}
