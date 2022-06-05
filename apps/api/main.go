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
	fmt.Println("PORT => " + port)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		enterprises := v1.Group("enterprises")
		{
			controller := new(controllers.EnterpriseController)
			// Enterprises
			enterprises.GET("/", controller.FindAllAsync)
			enterprises.GET("/:id", controller.FindByIdAsync)
			enterprises.POST("/", controller.CreateAsync)
			enterprises.PUT("/:id", controller.UpdateAsync)
			enterprises.DELETE("/:id", controller.DeleteAsync)
		}
	}

	router.Run(":" + port)
}
