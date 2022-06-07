package controllers

import (
	"net/http"

	enterpriseUseCase "acetime.com.br/business-crm/apps/api/application/use_cases"
	"acetime.com.br/business-crm/apps/api/domain"
	"github.com/gin-gonic/gin"
)

type EnterpriseController struct {
}

// @BasePath /api/v1

// FindAllEnterprises godoc
// @Summary Enterprises
// @Schemes
// @Description Find All Enterprises
// @Tags enterprise
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /enterprises [get]
func (e EnterpriseController) FindAll(c *gin.Context) {

	enterprises, err := enterpriseUseCase.FindAll()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, enterprises)

}

// FindEnterpriseById godoc
// @Summary Enterprise by id
// @Schemes
// @Description Find All Enterprises
// @Tags enterprise
// @Accept json
// @Produce json
// @Param id path string true "Enterprise ID"
// @Success 200 {string} Helloworld
// @Router /enterprises/{id} [get]
func (e EnterpriseController) FindById(c *gin.Context) {

	enterpriseId := c.Params.ByName("id")

	var enterprise domain.Enterprise

	enterprise, err := enterpriseUseCase.FindOneById(enterpriseId)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, enterprise)
}

// CreateEnterprise godoc
// @Summary					Create Enterprise
// @Schemes
// @Description 			Create Enterprise
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Success 				200 {string} Helloworld
// @Router 					/enterprises [post]
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

// UpdateEnterprise godoc
// @Summary 				Update Enterprise
// @Schemes
// @Description 			Update Enterprise
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Param					id path string true "EnterpriseID"
// @Success 				200 {string} Helloworld
// @Router 					/enterprises/{id} [put]
func (e EnterpriseController) Update(c *gin.Context) {

	var enterpriseId string
	var enterprise domain.Enterprise
	var err error

	enterpriseId = c.Params.ByName("id")

	c.BindJSON(&enterprise)

	err = enterpriseUseCase.Update(enterpriseId, enterprise)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DeleteEnterprise godoc
// @Summary 				Delete Enterprise
// @Schemes
// @Description 			Delete Enterprise
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Param					id path string true "EnterpriseID"
// @Success 				200 {string} Helloworld
// @Router 					/enterprises/{id} [delete]
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
