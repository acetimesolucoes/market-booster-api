package main

import (
	"fmt"
	"log"
	"os"

	"acetime.com.br/business-crm/apps/api/application/use_cases"
	"acetime.com.br/business-crm/apps/api/framework/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting the application...")
	utils.ConnectDB()

	err := godotenv.Load("apps/api/.env")

	if err == nil {
		log.Fatal("Failed with load .env file.")
	}

	port := os.Getenv("PORT")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		enterprises := v1.Group("enterprises")
		{
			// Enterprises
			enterprises.GET("/", use_cases.GetEnterprises)
			enterprises.GET("/:id", use_cases.GetEnterpriseById)
			enterprises.POST("/", use_cases.CreateEnterprise)
			enterprises.PUT("/:id", use_cases.UpdateEnterprise)
			enterprises.DELETE("/:id", use_cases.DeleteEnterprise)
		}
	}

	router.Run(":" + port)
}
