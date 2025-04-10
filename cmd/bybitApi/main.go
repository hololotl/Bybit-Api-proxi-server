package main

import (
	"bybit_api_servic_grpc/internal/app"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
	"golang.org/x/net/context"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	log := setupLogger(envLocal)
	conString := "user=postgres password=2005 dbname=petProjBybit sslmode=disable"
	application := app.NewApp(log, 8083, conString)
	application.GRPCServer.MustRun()
}

func GetTransaction() {
	client := bybit.NewBybitHttpClient("3w3LwpWhSbr9wIzKHA", "KRuqEkXwo681tWQou6rDEn32WMBvOvY4Hm0U", bybit.WithBaseURL(bybit.MAINNET))
	params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetTransactionLog(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
