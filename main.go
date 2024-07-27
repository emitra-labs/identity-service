package main

import (
	"context"
	"os"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/emitra-labs/common/amqp"
	"github.com/emitra-labs/common/mail"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/rest"
)

var port = env.GetIntDefault("PORT", 3000)

func init() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	db.Open()
	mail.Open(os.Getenv("SMTP_URL"))
}

func main() {
	s := rest.NewServer()

	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return s.Start(port)
	})

	m.AddShutdownJob(func() error {
		return s.Shutdown()
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
