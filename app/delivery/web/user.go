package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"net/http"
)

type UserDelivery struct {
	useCase domain.UserUseCase
}

func NewUserDelivery(useCase domain.UserUseCase) *UserDelivery {
	return &UserDelivery{useCase: useCase}
}

func (u *UserDelivery) CreateUser(c *gin.Context) {
	var userDto domain.UserDto
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := u.useCase.CreateUser(userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserDelivery) ActivateAccount(c *gin.Context) {
	activationPath := c.Param("activationpath")
	err := u.useCase.ActivateAccount(activationPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "account was activated successfully")
}
