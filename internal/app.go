package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-sousa-dev/meet-people/app/infra"
	"github.com/mateus-sousa-dev/meet-people/app/middlewares"
	docs "github.com/mateus-sousa-dev/meet-people/docs"
	"github.com/mateus-sousa-dev/meet-people/internal/emails"
	"github.com/mateus-sousa-dev/meet-people/internal/events"
	"github.com/mateus-sousa-dev/meet-people/internal/login"
	"github.com/mateus-sousa-dev/meet-people/internal/passwordresetconfigs"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := infra.StartDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	rabbitmqChannel, err := infra.StartRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	smtpDialer, err := infra.StartSmtpDialer()
	if err != nil {
		log.Fatal(err)
	}
	emailRepository := emails.NewRepository(smtpDialer)
	sendUseCase := emails.NewSendUseCase(emailRepository)
	emailDelivery := emails.NewDelivery(rabbitmqChannel, sendUseCase)
	go func() {
		err = emailDelivery.StartConsume()
		if err != nil {
			log.Fatal(err)
		}
	}()
	userRepository := users.NewRepository(db)
	eventRepository := events.NewRepository(rabbitmqChannel)
	passwordResetConfigRepo := passwordresetconfigs.NewRepository(db)
	writingUseCase := users.NewWritingUseCase(userRepository, eventRepository, passwordResetConfigRepo)
	userDelivery := users.NewDelivery(writingUseCase)
	loginUseCase := login.NewLoginUseCase(userRepository)
	loginDelivery := login.NewDelivery(loginUseCase)
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
	apiV1Routes.GET("/logged", middlewares.Authenticate(userDelivery.Logged))
	r.Run()
}
