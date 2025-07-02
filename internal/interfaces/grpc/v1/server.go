package grpc

import (
	userv1 "github.com/smile-ko/go-ddd-template/api/proto/user/v1/gen"
	"github.com/smile-ko/go-ddd-template/pkg/logger"
	"google.golang.org/grpc"
)

func RegisterGRPCV1Services(app *grpc.Server, log logger.ILogger) {
	h := NewUserHandler(log)
	userv1.RegisterUserServiceServer(app, h)
}
