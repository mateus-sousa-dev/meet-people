package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/internal/auth"
	"net/http"
)

type Delivery struct {
	writingUseCase WritingUseCase
}

func NewDelivery(writingUseCase WritingUseCase) *Delivery {
	return &Delivery{writingUseCase: writingUseCase}
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
func (u *Delivery) CreateUser(c *gin.Context) {
	var userDto UserDto
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := u.writingUseCase.CreateUser(userDto)
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
func (u *Delivery) ActivateAccount(c *gin.Context) {
	activationPath := c.Param("activationpath")
	err := u.writingUseCase.ActivateAccount(activationPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "account was activated successfully")
}

func (u *Delivery) Logged(c *gin.Context) {
	userID, err := auth.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(200, userID)
}

func (u *Delivery) ForgotPassword(c *gin.Context) {
	email := c.Query("email")
	err := u.writingUseCase.ForgotPassword(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url to reset password was send to your e-mail")
}

func (u *Delivery) ValidateUrlPassword(c *gin.Context) {
	url := c.Param("urlpasswordreset")
	err := u.writingUseCase.ValidateUrlPassword(url)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url is valid")
}

func (u *Delivery) ResetForgottenPassword(c *gin.Context) {
	url := c.Param("urlpasswordreset")
	var passwordDto PasswordDto
	err := c.ShouldBindJSON(&passwordDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = u.writingUseCase.ResetForgottenPassword(passwordDto, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "password updated successfully")
}
