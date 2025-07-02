package grpc

import (
	"context"

	userv1 "github.com/smile-ko/go-ddd-template/api/proto/user/v1/gen"
	"github.com/smile-ko/go-ddd-template/pkg/logger"
)

type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	log logger.ILogger
}

func NewUserHandler(log logger.ILogger) *UserHandler {
	return &UserHandler{
		log: log,
	}
}

func (h *UserHandler) GetUserById(ctx context.Context, req *userv1.GetUserByIdReq) (*userv1.PublicUserInfoResp, error) {
	return nil, nil
}
