package controllers

import (
	"net/http"

	enterpriseUseCase "acetime.com.br/business-crm/apps/api/application/use_cases"
	"github.com/gin-gonic/gin"
)

type EnterpriseController struct {
}

func (e EnterpriseController) FindAll(c *gin.Context) {

	enterprises, err := enterpriseUseCase.Find()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, enterprises)

}

func (e EnterpriseController) FindById(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, []string{})
}

func (e EnterpriseController) Create(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, nil)
}

func (e EnterpriseController) Update(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}

func (e EnterpriseController) Delete(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}
