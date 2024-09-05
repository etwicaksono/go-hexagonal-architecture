package grpc

import (
	"context"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) SendMultimediaMessage(context.Context, *example.SendMultimediaMessageRequest) (*emptypb.Empty, error) {
	// TODO: implement this
	panic("unimplemented")
}
