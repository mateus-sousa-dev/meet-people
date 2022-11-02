package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Delivery struct {
	loginUseCase LoginUseCase
}

func NewDelivery(loginUseCase LoginUseCase) *Delivery {
	return &Delivery{loginUseCase: loginUseCase}
}

func (l *Delivery) Exec(c *gin.Context) {
	var login LoginDto
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
