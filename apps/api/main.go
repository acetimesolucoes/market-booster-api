package main

import (
	"acetime-business-crm/apps/api/application/use_cases"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Enterprises
		v1.GET("/enterprises", use_cases.GetEnterprises)
		v1.GET("/enterprises/:id", use_cases.GetEnterpriseById)
		v1.POST("/enterprises", use_cases.CreateEnterprise)
		v1.PUT("/enterprises/:id", use_cases.UpdateEnterprise)
		v1.DELETE("/enterprises/:id", use_cases.DeleteEnterprise)
	}

	router.Run("localhost:5000")
}
