package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"marketbooster/docs"
	"marketbooster/framework/http/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Acetime Business API Documentation
// @version 1.0.0

// @contact.name Acetime Soluções
// @contact.email develop@acetime.com.br
// @contact.url https://www.acetime.com.br

// @host localhost:3000
// @basePath /api/v1

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars %v", err)
	}

	port := os.Getenv("PORT")

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// @BasePath /api/v1
	v1 := router.Group("/api/v1")
	{
		// Authentication router
		new(routers.AuthenticationRouter).CreateRouter(v1)
		// Enterprises router
		new(routers.EnterpriseRouter).CreateRouter(v1)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	fmt.Printf("Server is listner in port %s", port)

	router.Run(":" + port)
}
