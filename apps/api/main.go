package main

import (
	"fmt"
	"log"
	"os"

	"acetime.com.br/business-crm/apps/api/docs"
	"acetime.com.br/business-crm/apps/api/framework/http/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Acetime Business API Documentation
// @version 1.0.0

// @host localhost:3000
// @basePath /api/v1

func main() {

	// Loading .env file
	err := godotenv.Load("env")
	if err == nil {
		fmt.Println(err)
		log.Fatal("Failed with load .env file.")
	}

	port := os.Getenv("PORT")

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// @BasePath /api/v1
	v1 := router.Group("/api/v1")
	{
		// todo: moove the router to external file
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":" + port)
}
