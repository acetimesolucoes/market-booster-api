package controllers

import (
	"net/http"

	enterpriseUseCase "acetime.com.br/business-crm/apps/api/application/use_cases"
	"acetime.com.br/business-crm/apps/api/domain"
	"github.com/gin-gonic/gin"
)

type EnterpriseController struct {
}

func (e EnterpriseController) FindAll(c *gin.Context) {

	enterprises, err := enterpriseUseCase.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, enterprises)

}

func (e EnterpriseController) FindById(c *gin.Context) {

	enterpriseId := c.Params.ByName("id")

	var enterprise domain.Enterprise

	enterprise, err := enterpriseUseCase.FindOneById(enterpriseId)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, enterprise)
}

func (e EnterpriseController) Create(c *gin.Context) {

	var err error
	var enterprise domain.Enterprise

	err = c.BindJSON(&enterprise)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = enterpriseUseCase.Create(enterprise)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusCreated, nil)

}

func (e EnterpriseController) Update(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}

func (e EnterpriseController) Delete(c *gin.Context) {

	var enterpriseId string
	var err error

	enterpriseId = c.Params.ByName("id")

	err = enterpriseUseCase.Delete(enterpriseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}
