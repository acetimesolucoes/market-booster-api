package controllers

import (
	"net/http"
	"strconv"

	userUseCase "marketbooster/application/use_cases/user_use_cases"
	"marketbooster/domain"
	http_exception "marketbooster/framework/exception"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// @BasePath /api/v1

// FindAllUsers godoc
// @Summary 			Users
// @Schemes
// @Description 		Find All Users
// @Tags 				user
// @Accept 				json
// @Produce 			json
// @Param 				page query int false "Current page to paginate"
// @Router /users [get]
func (e *UserController) FindAll(c *gin.Context) {

	page, pageErr := strconv.ParseInt(c.Query("page"), 0, 64)
	limit, limitErr := strconv.ParseInt(c.Query("limit"), 0, 64)

	if pageErr != nil {
		page = 1
	}

	if limitErr != nil {
		limit = 25
	}

	users, err := userUseCase.FindAll(page, limit)

	result := new(http_exception.HttpSuccess[domain.Users])

	result.Data = users
	result.Page = page
	result.Limit = limit

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, result)

}

// FindUserById godoc
// @Summary User by id
// @Schemes
// @Description Find All Users
// @Tags enterprise
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} Helloworld
// @Router /users/{id} [get]
func (e *UserController) FindById(c *gin.Context) {

	enterpriseId := c.Params.ByName("id")

	var enterprise domain.User

	enterprise, err := enterpriseUseCase.FindOneById(enterpriseId)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, enterprise)
}

// CreateUser godoc
// @Summary					Create User
// @Schemes
// @Description 			Create User
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Success 				200 {string} Helloworld
// @Router 					/users [post]
func (e *UserController) Create(c *gin.Context) {

	var err error
	var enterprise domain.User

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

// UpdateUser godoc
// @Summary 				Update User
// @Schemes
// @Description 			Update User
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Param					id path string true "UserID"
// @Success 				200 {string} Helloworld
// @Router 					/users/{id} [put]
func (e *UserController) Update(c *gin.Context) {

	var enterpriseId string
	var enterprise domain.User
	var err error

	enterpriseId = c.Params.ByName("id")

	c.BindJSON(&enterprise)

	err = enterpriseUseCase.Update(enterpriseId, enterprise)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DeleteUser godoc
// @Summary 				Delete User
// @Schemes
// @Description 			Delete User
// @Tags enterprise
// @Accept 					json
// @Produce 				json
// @Param					id path string true "UserID"
// @Success 				200 {string} Helloworld
// @Router 					/users/{id} [delete]
func (e *UserController) Delete(c *gin.Context) {

	var enterpriseId string
	var err error

	enterpriseId = c.Params.ByName("id")

	err = enterpriseUseCase.Delete(enterpriseId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}
