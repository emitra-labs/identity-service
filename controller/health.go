package controller

import (
	"context"

	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/identity-service/model"
)

func HealthCheck(ctx context.Context, req *commonModel.Empty) (*model.Health, error) {
	return &model.Health{Status: "ok"}, nil
}
