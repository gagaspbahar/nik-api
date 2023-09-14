package util

import "github.com/gin-gonic/gin"

func MakeResponse(status int, message string, err error, data any) gin.H {
	if err != nil {
		return gin.H{
			"status":  status,
			"message": message,
			"data": gin.H{
				"error": err.Error(),
			},
		}
	}
	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func MakeMultipleErrorResponse(status int, message string, errors []error) gin.H {
	var errorMessages []string
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}
	return gin.H{
		"status":  status,
		"message": message,
		"data": gin.H{
			"errors": errorMessages,
		},
	}
}
