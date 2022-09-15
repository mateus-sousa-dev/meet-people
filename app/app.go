package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-sousa-dev/meet-people/app/delivery/rabbitmq"
	"github.com/mateus-sousa-dev/meet-people/app/delivery/web"
	"github.com/mateus-sousa-dev/meet-people/app/infra"
	"github.com/mateus-sousa-dev/meet-people/app/middlewares"
	"github.com/mateus-sousa-dev/meet-people/app/repository"
	"github.com/mateus-sousa-dev/meet-people/app/usecase"
	docs "github.com/mateus-sousa-dev/meet-people/docs"
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
	mailRepository := repository.NewMailRepository(smtpDialer)
	emailUsecase := usecase.NewEmailUseCase(mailRepository)
	emailDelivery := rabbitmq.NewEmailDelivery(rabbitmqChannel, emailUsecase)
	go func() {
		err = emailDelivery.StartConsume()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		err = emailDelivery.StartConsume2()
		if err != nil {
			log.Fatal(err)
		}
	}()
	userRepository := repository.NewUserRepository(db)
	eventRepository := repository.NewEventRepository(rabbitmqChannel)
	userUseCase := usecase.NewUserUseCase(userRepository, eventRepository)
	userDelivery := web.NewUserDelivery(userUseCase)
	loginUseCase := usecase.NewLoginUseCase(userRepository)
	loginDelivery := web.NewLoginDelivery(loginUseCase)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	apiV1Routes := r.Group("/api/v1")
	apiV1Routes.POST("/users", userDelivery.CreateUser)
	apiV1Routes.GET("/activate-account/:activationpath", userDelivery.ActivateAccount)
	apiV1Routes.PATCH("/forgot-password", userDelivery.ForgotPassword)
	apiV1Routes.POST("/login", loginDelivery.Exec)
	apiV1Routes.GET("/logged", middlewares.Authenticate(userDelivery.Logged))
	r.Run()
}
