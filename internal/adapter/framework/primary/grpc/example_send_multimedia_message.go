package grpc

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/utils/error_util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) SendMultimediaMessage(ctx context.Context, request *example.SendMultimediaMessageRequest) (*emptypb.Empty, error) {
	var files []entity.MultimediaFile
	for _, file := range request.Files {
		files = append(files, entity.MultimediaFile{
			Filename:    file.Filename,
			ContentType: file.ContentType,
			Data:        file.Data,
		})
	}
	err := a.handler.ExampleMessageApp.SendMultimediaMessage(ctx, entity.SendMultimediaMessageRequest{
		Sender:   request.Sender,
		Receiver: request.Receiver,
		Message:  request.Message,
		Storage:  entity.MultimediaStorage(request.Storage),
		Files:    files,
	})
	if err != nil {
		if customError, isCustomError := error_util.IsCustomError(err); isCustomError {
			for k, v := range customError.Fields {
				return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s: %s", k, v))
			}
		}
		slog.ErrorContext(ctx, "Failed to send text message", slog.String(entity.Error, err.Error()))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
