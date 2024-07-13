package main

import (
	"context"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/rest"
)

var port = env.GetIntDefault("PORT", 3000)

func init() {
	db.Open()
}

func main() {
	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return rest.Server.Start(port)
	})

	m.AddShutdownJob(func() error {
		return rest.Server.Shutdown()
	})

	<-m.Done()
}
