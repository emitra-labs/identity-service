package auth

import (
	"context"

	commonConstant "github.com/ukasyah-dev/common/constant"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/identity-service/controller/session"
	"github.com/ukasyah-dev/identity-service/model"
)

func SignOut(ctx context.Context, req *commonModel.Empty) (*commonModel.BasicResponse, error) {
	sessionID, _ := ctx.Value(commonConstant.SessionID).(string)

	_, err := session.DeleteSession(ctx, &model.DeleteSessionRequest{
		ID: sessionID,
	})
	if err != nil {
		return nil, err
	}

	return &commonModel.BasicResponse{
		Message: "User sign-out was successful",
	}, nil
}
