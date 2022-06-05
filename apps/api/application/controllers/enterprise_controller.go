package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnterpriseController struct {
}

func (e EnterpriseController) FindAllAsync(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, []string{})
}

func (e EnterpriseController) FindByIdAsync(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, []string{})
}

func (e EnterpriseController) CreateAsync(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, nil)
}

func (e EnterpriseController) UpdateAsync(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}

func (e EnterpriseController) DeleteAsync(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, nil)
}
