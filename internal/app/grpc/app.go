package grpc

import (
	authgrpc "bybit_api_servic_grpc/internal/grpc/bybitApi"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log"
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
		log.Fatal(err)
	}
}

func (app *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", err)
	}
	if err := app.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", err)
	}
	return nil

}
