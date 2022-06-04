package use_cases

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEnterprises(c *gin.Context) {
	fmt.Println("Start to get enterprises")
	c.IndentedJSON(http.StatusOK, []string{})
}

func GetEnterpriseById(c *gin.Context) {
	fmt.Println("Start to get enterprises")
	c.IndentedJSON(http.StatusOK, []string{})
}

func CreateEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{})
}

func UpdateEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{})
}

func DeleteEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []string{})
}
