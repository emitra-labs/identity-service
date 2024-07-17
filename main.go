package main

import (
	"context"
	"os"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/ukasyah-dev/common/amqp"
	"github.com/ukasyah-dev/common/mail"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/rest"
)

var port = env.GetIntDefault("PORT", 3000)

func init() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	db.Open()
	mail.Open(os.Getenv("SMTP_URL"))
}

func main() {
	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return rest.Server.Start(port)
	})

	m.AddShutdownJob(func() error {
		return rest.Server.Shutdown()
	})

	m.AddShutdownJob(func() error {
		return mail.Close()
	})

	m.AddShutdownJob(func() error {
		return db.Close()
	})

	m.AddShutdownJob(func() error {
		return amqp.Close()
	})

	<-m.Done()
}
