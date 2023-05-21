package controllers

import (
	"net/http"
	"strconv"

	"marketbooster/application/use_cases"
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

	users, err := new(use_cases.UserUseCase).FindAll(page, limit)

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
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} Helloworld
// @Router /users/{id} [get]
func (e *UserController) FindById(c *gin.Context) {

	userId := c.Params.ByName("id")

	var user domain.User

	user, err := new(use_cases.UserUseCase).FindOneById(userId)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary					Create User
// @Schemes
// @Description 			Create User
// @Tags user
// @Accept 					json
// @Produce 				json
// @Success 				200 {string} Helloworld
// @Router 					/users [post]
func (e *UserController) Create(c *gin.Context) {

	var err error
	var user domain.User

	err = c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err = new(use_cases.UserUseCase).Create(user)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusCreated, nil)

}

// UpdateUser godoc
// @Summary 				Update User
// @Schemes
// @Description 			Update User
// @Tags user
// @Accept 					json
// @Produce 				json
// @Param					id path string true "UserID"
// @Success 				200 {string} Helloworld
// @Router 					/users/{id} [put]
func (e *UserController) Update(c *gin.Context) {

	var userId string
	var user domain.User
	var err error

	userId = c.Params.ByName("id")

	c.BindJSON(&user)

	err = new(use_cases.UserUseCase).Update(userId, user)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// DeleteUser godoc
// @Summary 				Delete User
// @Schemes
// @Description 			Delete User
// @Tags user
// @Accept 					json
// @Produce 				json
// @Param					id path string true "UserID"
// @Success 				200 {string} Helloworld
// @Router 					/users/{id} [delete]
func (e *UserController) Delete(c *gin.Context) {

	var userId string
	var err error

	userId = c.Params.ByName("id")

	err = new(use_cases.UserUseCase).Delete(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, nil)
}
