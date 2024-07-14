package controller

import (
	"context"

	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/identity-service/model"
)

func HealthCheck(ctx context.Context, req *commonModel.Empty) (*model.Health, error) {
	return &model.Health{Status: "ok"}, nil
}
