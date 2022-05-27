package use_cases

import (
	"acetime-business-crm/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var enterprises = []domain.Enterprise{
	{ID: "1", BusinessName: "Seegal Comercio do Brasil LTDA"},
	{ID: "2", BusinessName: "Acetime Soluções Digitais LTDA"},
}

func getEnterprises(c *gin.Context) {
	fmt.Println("Start to get enterprises")
	c.IndentedJSON(http.StatusOK, enterprises)
}

func createEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}

func updateEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}

func deleteEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}
