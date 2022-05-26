package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Enterprise struct {
	ID                 string      `json:"_id"`
	LegalIdentifier    string      `json:"legalIdentifier"`
	LogoUrl            string      `json:"LogoUrl"`
	IsFilial           string      `json:"isFilial"`
	BusinessName       string      `json:"businessName"`
	FantasyName        string      `json:"fantasyName"`
	RegistrationStatus string      `json:"registrationStatus"`
	LegalNature        string      `json:"legalNature"`
	CNAEPrincipal      string      `json:"CNAEPrincipal"`
	CNAESecondary      string      `json:"CNAESecondary"`
	Users              []string    `json:"users"`
	Customers          []Customer  `json:"customers"`
	Employeers         []Employeer `json:"employeers"`
}

type Customer struct {
	ID              string `json:"_id"`
	Email           string
	EnterpriseId    string
	FullName        string
	LegalNature     string
	LegalIdentifier string
	Phones          []string
	CreatedDate     time.Time
	CreatedBy       string
	UpdatedDate     time.Time
	UpdatedBy       string
}

type Employeer struct {
}

var enterprises = []Enterprise{
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

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Enterprises
		v1.GET("/enterprises", getEnterprises)
		v1.POST("/enterprises", createEnterprise)
		v1.PUT("/enterprises/:id", updateEnterprise)
		v1.DELETE("/enterprises/:id", deleteEnterprise)
	}

	router.Run("localhost:5000")
}
