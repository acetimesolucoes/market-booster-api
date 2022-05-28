package use_cases

import (
	"fmt"
	"net/http"

	"acetime.com.br/business-crm/apps/api/domain"

	"github.com/gin-gonic/gin"
)

var enterprises = []*domain.Enterprise{
	{
		ID:                 "1",
		BusinessName:       "Seegal Comercio do Brasil LTDA",
		LegalIdentifier:    "",
		LogoUrl:            "",
		IsFilial:           false,
		FantasyName:        "",
		RegistrationStatus: "",
		LegalNature:        "",
		CNAEPrincipal:      "",
		CNAESecondary:      "",
		Users:              []string{},
		Customers:          []domain.Customer{},
		Employeers:         []domain.Employeer{},
	},
}

func GetEnterprises(c *gin.Context) {
	fmt.Println("Start to get enterprises")
	c.IndentedJSON(http.StatusOK, enterprises)
}

func GetEnterpriseById(c *gin.Context) {
	fmt.Println("Start to get enterprises")
	c.IndentedJSON(http.StatusOK, enterprises)
}

func CreateEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}

func UpdateEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}

func DeleteEnterprise(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, enterprises)
}
