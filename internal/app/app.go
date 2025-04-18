package app

import (
	grpcapp "bybit_api_servic_grpc/internal/app/grpc"
	callApi2 "bybit_api_servic_grpc/internal/callApi"
	"bybit_api_servic_grpc/internal/services/apiBybit"
	postgr "bybit_api_servic_grpc/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string) *App {
	const op = "NewApp"
	log.With("port", grpcPort, "storagePath", storagePath, "op", op).Info("NewApp")
	storage, err := postgr.NewStorage(storagePath)
	if err != nil {
		log.With("storagePath", storagePath, "op", op).Error("NewApp: failed to open storage")
	}
	callApi := callApi2.NewCallApi(log)
	authServer := apiBybit.NewBybit(storage, callApi, log)
	grpcApp := grpcapp.NewApp(log, grpcPort, authServer)
	return &App{GRPCServer: grpcApp}
}
