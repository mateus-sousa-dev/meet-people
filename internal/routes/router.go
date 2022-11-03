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

func SetupRoutes(userDelivery users.Delivery, loginDelivery login.Delivery, friendshipDelivery friendships.Delivery) {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiV1Routes := r.Group("/api/v1")
	apiV1Routes.POST("/users", userDelivery.CreateUser)
	apiV1Routes.GET("/activate-account/:activationpath", userDelivery.ActivateAccount)
	apiV1Routes.PATCH("/forgot-password", userDelivery.ForgotPassword)
	apiV1Routes.GET("/validate-url-password/:urlpasswordreset", userDelivery.ValidateUrlPassword)
	apiV1Routes.PATCH("/reset-forgotten-password/:urlpasswordreset", userDelivery.ResetForgottenPassword)
	apiV1Routes.POST("/login", loginDelivery.Exec)
	apiV1Routes.POST("/friendship", middlewares.Authenticate(friendshipDelivery.RequestFriendship))
	apiV1Routes.GET("/logged", middlewares.Authenticate(userDelivery.Logged))
	r.Run()
}
