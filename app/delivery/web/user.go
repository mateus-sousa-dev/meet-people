package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/app/auth"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"net/http"
)

type UserDelivery struct {
	useCase domain.UserUseCase
}

func NewUserDelivery(useCase domain.UserUseCase) *UserDelivery {
	return &UserDelivery{useCase: useCase}
}

// CreateUser godoc
// @Summary Create a user
// @Description Route to create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.UserDto true "Modelo de usuário"
// @Success 201 {object} domain.User
// @Router /api/v1/users [post]
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
	c.JSON(http.StatusCreated, user)
}

// ActivateAccount godoc
// @Summary Activate Account
// @Description Route to Activate Account
// @Tags Users
// @Accept json
// @Produce json
// @Param activationpath query string true "Path de ativação"
// @Success 201 {string} account was activated successfully
// @Router /api/v1/activate-account/{activationpath} [get]
func (u *UserDelivery) ActivateAccount(c *gin.Context) {
	activationPath := c.Param("activationpath")
	err := u.useCase.ActivateAccount(activationPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "account was activated successfully")
}

func (u *UserDelivery) Logged(c *gin.Context) {
	userID, err := auth.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(200, userID)
}

func (u *UserDelivery) ForgotPassword(c *gin.Context) {
	email := c.Query("email")
	err := u.useCase.ForgotPassword(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url to reset password was send to your e-mail")
}

func (u *UserDelivery) ValidateUrlPassword(c *gin.Context) {
	url := c.Param("urlpasswordreset")
	err := u.useCase.ValidateUrlPassword(url)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url is valid")
}

func (u *UserDelivery) ResetForgottenPassword(c *gin.Context) {
	url := c.Param("urlpasswordreset")
	var passwordDto domain.PasswordDto
	err := c.ShouldBindJSON(&passwordDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = u.useCase.ResetForgottenPassword(passwordDto, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "password updated successfully")
}
