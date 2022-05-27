package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Enterprises
		v1.GET("/enterprises", use_cases.getEnterprises)
		v1.POST("/enterprises", use_cases.createEnterprise)
		v1.PUT("/enterprises/:id", use_cases.updateEnterprise)
		v1.DELETE("/enterprises/:id", use_cases.deleteEnterprise)
	}

	router.Run("localhost:5000")
}
