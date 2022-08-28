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
	c.JSON(http.StatusOK, "OI")
}
