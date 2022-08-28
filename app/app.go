package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mateus-sousa-dev/meet-people/app/delivery/web"
	"github.com/mateus-sousa-dev/meet-people/app/infra"
	"github.com/mateus-sousa-dev/meet-people/app/repository"
	"github.com/mateus-sousa-dev/meet-people/app/usecase"
	"log"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := infra.StartConnection()
	if err != nil {
		log.Fatal(err)
	}
	smtpDialer, err := infra.StartSmtpDialer()
	if err != nil {
		log.Fatal(err)
	}
	mailRepository := repository.NewMailRepository(smtpDialer)
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, mailRepository)
	userDelivery := web.NewUserDelivery(userUseCase)
	r := gin.Default()
	apiV1Routes := r.Group("/api/v1")
	apiV1Routes.POST("/users", userDelivery.CreateUser)
	r.Run()
}
