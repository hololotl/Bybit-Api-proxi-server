package grpc

import (
	authgrpc "bybit_api_servic_grpc/internal/grpc/bybitApi"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int, authServer authgrpc.ApiFunctions) *App {
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(recovery.UnaryServerInterceptor()))
	authgrpc.RegisterApiFunctions(grpcServer, authServer)
	return &App{log: log, grpcServer: grpcServer, port: port}
}

func (app *App) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func (app *App) Run() error {
	const op = "Run"

	app.log.With("port", app.port).Info("Starting server")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if err := app.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	
	return nil

}
