package controller

import (
	"context"
	"fmt"

	"github.com/ukasyah-dev/identity-service/model"
)

func Hello(ctx context.Context, req *model.HelloRequest) (*model.HelloResponse, error) {
	message := "Hello, stranger!"
	if req != nil && req.Name != "" {
		message = fmt.Sprintf("Hi, %s!", req.Name)
	}

	return &model.HelloResponse{Message: message}, nil
}
