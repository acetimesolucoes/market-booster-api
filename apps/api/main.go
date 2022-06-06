package main

import (
	"fmt"
	"log"
	"os"

	"acetime.com.br/business-crm/apps/api/application/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Loading .env file
	err := godotenv.Load("env")
	if err == nil {
		fmt.Println(err)
		log.Fatal("Failed with load .env file.")
	}

	port := os.Getenv("PORT")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		enterprises := v1.Group("enterprises")
		{
			enterpriseController := new(controllers.EnterpriseController)
			// Enterprises
			enterprises.GET("/", enterpriseController.FindAll)
			enterprises.GET("/:id", enterpriseController.FindById)
			enterprises.POST("/", enterpriseController.Create)
			enterprises.PUT("/:id", enterpriseController.Update)
			enterprises.DELETE("/:id", enterpriseController.Delete)
		}
	}

	router.Run(":" + port)
}
