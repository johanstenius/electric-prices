package v1

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("hello-world", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello World"})
		})
	}

	router.Run(":8080")
}
