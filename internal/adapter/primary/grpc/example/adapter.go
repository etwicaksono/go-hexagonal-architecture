package example

import (
	"context"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/ports/app"
	primaryPort "github.com/etwicaksono/go-hexagonal-architecture/internal/ports/primary/grpc"
	"github.com/etwicaksono/public-proto/gen/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type Adapter struct {
	ctx        context.Context
	address    string
	exampleApp app.ExampleAppInterface
}

func NewExampleGrpcAdapter(
	ctx context.Context,
	address string,
	exampleApp app.ExampleAppInterface,
) primaryPort.ExampleGrpcHandlerInterface {
	return &Adapter{
		ctx:        ctx,
		address:    address,
		exampleApp: exampleApp,
	}
}

func (adapter *Adapter) Run() error {
	// Initialize net listener
	listen, err := net.Listen("tcp", adapter.address)
	if err != nil {
		slog.ErrorContext(adapter.ctx, "Failed to listen on port", slog.String("address", adapter.address), slog.String(entity.Error, err.Error()))
		return err
	}

	// Initialize grpc server
	grpcServer := grpc.NewServer()
	example.RegisterExampleServiceServer(grpcServer, NewExampleGrpcAdapter(adapter.ctx, adapter.address, adapter))
	slog.InfoContext(adapter.ctx, "grpc server running ", slog.String("address", adapter.address))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Run grpc server
	err = grpcServer.Serve(listen)
	if err != nil {
		slog.WarnContext(adapter.ctx, "Failed to serve grpc server", slog.String(entity.Error, err.Error()))
		return err
	}

	slog.InfoContext(adapter.ctx, "grpc server stopped")
	return nil
}
