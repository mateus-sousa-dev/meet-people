package internal

import (
	"github.com/joho/godotenv"
	"github.com/mateus-sousa-dev/meet-people/internal/emails"
	"github.com/mateus-sousa-dev/meet-people/internal/events"
	"github.com/mateus-sousa-dev/meet-people/internal/friendships"
	"github.com/mateus-sousa-dev/meet-people/internal/infra"
	"github.com/mateus-sousa-dev/meet-people/internal/login"
	"github.com/mateus-sousa-dev/meet-people/internal/passwordresetconfigs"
	"github.com/mateus-sousa-dev/meet-people/internal/routes"
	"github.com/mateus-sousa-dev/meet-people/internal/users"
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
	readingUseCase := users.NewReadingUseCase(userRepository)
	userDelivery := users.NewDelivery(writingUseCase, readingUseCase)
	loginUseCase := login.NewLoginUseCase(userRepository)
	loginDelivery := login.NewDelivery(loginUseCase)
	friendshipRepo := friendships.NewRepository(db)
	friendshipWritingUseCase := friendships.NewWritingUseCase(friendshipRepo, userRepository)
	friendshipDelivery := friendships.NewDelivery(friendshipWritingUseCase)
	routes.SetupRoutes(routes.RouterDeliveriesDto{
		UserDelivery:       userDelivery,
		LoginDelivery:      loginDelivery,
		FriendshipDelivery: friendshipDelivery,
	})
}
