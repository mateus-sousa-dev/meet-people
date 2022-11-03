package friendships

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-sousa-dev/meet-people/internal/auth"
	"net/http"
)

type Delivery interface {
	RequestFriendship(c *gin.Context)
}

type delivery struct {
	writingUseCase WritingUseCase
}

func NewDelivery(writingUseCase WritingUseCase) Delivery {
	return &delivery{writingUseCase: writingUseCase}
}

// RequestFriendship godoc
// @Summary Request a friendship
// @Description Route to request a friendship
// @Tags Friendships
// @Accept json
// @Produce json
// @Param friendship body FriendshipDto true "Modelo de amizade"
// @Success 201 {object} Friendship
// @Router /api/v1/friendship [post]
func (d *delivery) RequestFriendship(c *gin.Context) {
	loggedUserID, err := auth.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	var friendshipDto FriendshipDto
	err = c.ShouldBindJSON(&friendshipDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	friendship, err := d.writingUseCase.RequestFriendship(friendshipDto, loggedUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, friendship)
}
