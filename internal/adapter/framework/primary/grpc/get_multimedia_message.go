package grpc

import (
	"context"

	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *adapter) GetMultimediaMessage(context.Context, *emptypb.Empty) (*example.GetMultimediaMessageResponse, error) {
	// TODO: implement this
	panic("unimplemented")
}
