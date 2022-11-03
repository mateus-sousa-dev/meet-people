package routes

import (
	"github.com/gin-gonic/gin"
	docs "github.com/mateus-sousa-dev/meet-people/docs"
	"github.com/mateus-sousa-dev/meet-people/internal/friendships"
	"github.com/mateus-sousa-dev/meet-people/internal/login"
	"github.com/mateus-sousa-dev/meet-people/internal/middlewares"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterDeliveriesDto struct {
	UserDelivery       users.Delivery
	LoginDelivery      login.Delivery
	FriendshipDelivery friendships.Delivery
}

func SetupRoutes(routerDeliveriesDto RouterDeliveriesDto) {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiV1Routes := r.Group("/api/v1")
	apiV1Routes.POST("/users", routerDeliveriesDto.UserDelivery.CreateUser)
	apiV1Routes.GET("/activate-account/:activationpath", routerDeliveriesDto.UserDelivery.ActivateAccount)
	apiV1Routes.PATCH("/forgot-password", routerDeliveriesDto.UserDelivery.ForgotPassword)
	apiV1Routes.GET("/validate-url-password/:urlpasswordreset", routerDeliveriesDto.UserDelivery.ValidateUrlPassword)
	apiV1Routes.PATCH("/reset-forgotten-password/:urlpasswordreset", routerDeliveriesDto.UserDelivery.ResetForgottenPassword)
	apiV1Routes.POST("/login", routerDeliveriesDto.LoginDelivery.Exec)
	apiV1Routes.POST("/friendship", middlewares.Authenticate(routerDeliveriesDto.FriendshipDelivery.RequestFriendship))
	apiV1Routes.PATCH("/friendship/:friendshipId", middlewares.Authenticate(routerDeliveriesDto.FriendshipDelivery.AcceptFriendship))
	r.Run()
}
