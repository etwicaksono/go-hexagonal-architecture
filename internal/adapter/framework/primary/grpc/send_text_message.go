package grpc

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	utils2 "github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) SendTextMessage(ctx context.Context, request *example.SendTextMessageRequest) (*emptypb.Empty, error) {
	err := a.handler.ExampleApp.SendTextMessage(ctx, entity.SendTextMessageRequest{
		Sender:   request.Sender,
		Receiver: request.Receiver,
		Message:  request.Message,
	})
	if err != nil {
		if customError, isCustomError := utils2.IsCustomError(err); isCustomError {
			for k, v := range customError.Fields {
				return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s: %s", k, v))
			}
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
