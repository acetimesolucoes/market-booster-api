package main

import (
	"fmt"
	"log"
	"os"

	"acetime.com.br/business-crm/apps/api/application/controllers"
	"acetime.com.br/business-crm/apps/api/framework/utils"

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

	// Connecting to database
	utils.ConnectDB()

	port := os.Getenv("PORT")

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		enterprises := v1.Group("enterprises")
		{
			controller := new(controllers.EnterpriseController)
			// Enterprises
			enterprises.GET("/", controller.FindAll)
			enterprises.GET("/:id", controller.FindById)
			enterprises.POST("/", controller.Create)
			enterprises.PUT("/:id", controller.Update)
			enterprises.DELETE("/:id", controller.Delete)
		}
	}

	router.Run(":" + port)
}
