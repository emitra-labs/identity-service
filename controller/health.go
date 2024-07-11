package controller

import (
	"context"

	"github.com/ukasyah-dev/identity-service/model"
)

func HealthCheck(ctx context.Context, req *model.Empty) (*model.Health, error) {
	return &model.Health{Status: "ok"}, nil
}
