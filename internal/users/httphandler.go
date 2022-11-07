package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/internal/auth"
	"net/http"
)

type Delivery interface {
	CreateUser(c *gin.Context)
	ActivateAccount(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ValidateUrlPassword(c *gin.Context)
	ResetForgottenPassword(c *gin.Context)
	GetMyFriends(c *gin.Context)
}

type delivery struct {
	writingUseCase WritingUseCase
	readingUseCase ReadingUseCase
}

func NewDelivery(writingUseCase WritingUseCase, readingUseCase ReadingUseCase) Delivery {
	return &delivery{writingUseCase: writingUseCase, readingUseCase: readingUseCase}
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
func (u *delivery) CreateUser(c *gin.Context) {
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
func (u *delivery) ActivateAccount(c *gin.Context) {
	activationPath := c.Param("activationpath")
	err := u.writingUseCase.ActivateAccount(activationPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "account was activated successfully")
}

func (u *delivery) ForgotPassword(c *gin.Context) {
	email := c.Query("email")
	err := u.writingUseCase.ForgotPassword(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url to reset password was send to your e-mail")
}

func (u *delivery) ValidateUrlPassword(c *gin.Context) {
	url := c.Param("urlpasswordreset")
	err := u.writingUseCase.ValidateUrlPassword(url)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, "url is valid")
}

func (u *delivery) ResetForgottenPassword(c *gin.Context) {
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

func (u *delivery) GetMyFriends(c *gin.Context) {
	loggedUserID, err := auth.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	users := u.readingUseCase.GetMyFriends(loggedUserID)
	c.JSON(http.StatusOK, users)
}
