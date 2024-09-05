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

type adapter struct {
	ctx        context.Context
	address    string
	exampleApp app.ExampleAppInterface
}

func NewExampleGrpcAdapter(
	ctx context.Context,
	address string,
	exampleApp app.ExampleAppInterface,
) primaryPort.ExampleGrpcHandlerInterface {
	return &adapter{
		ctx:        ctx,
		address:    address,
		exampleApp: exampleApp,
	}
}

func (a *adapter) Run() error {
	// Initialize net listener
	listen, err := net.Listen("tcp", a.address)
	if err != nil {
		slog.ErrorContext(
			a.ctx,
			"Failed to listen on port",
			slog.String("address", a.address),
			slog.String(entity.Error, err.Error()))
		return err
	}

	// Initialize grpc server
	grpcServer := grpc.NewServer()
	example.RegisterExampleServiceServer(
		grpcServer,
		NewExampleGrpcAdapter(
			a.ctx,
			a.address,
			a.exampleApp,
		),
	)
	slog.InfoContext(a.ctx, "grpc server running ", slog.String("address", a.address))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	// Run grpc server
	err = grpcServer.Serve(listen)
	if err != nil {
		slog.WarnContext(a.ctx, "Failed to serve grpc server", slog.String(entity.Error, err.Error()))
		return err
	}

	slog.InfoContext(a.ctx, "grpc server stopped")
	return nil
}
