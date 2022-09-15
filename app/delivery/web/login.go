package web

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/app/domain"
	"net/http"
)

type LoginDelivery struct {
	loginUseCase domain.LoginUseCase
}

func NewLoginDelivery(loginUseCase domain.LoginUseCase) *LoginDelivery {
	return &LoginDelivery{loginUseCase: loginUseCase}
}

func (l *LoginDelivery) Exec(c *gin.Context) {
	var login domain.LoginDto
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := l.loginUseCase.Exec(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
