package apiBybit

import (
	grpcFile "bybit_api_servic_grpc/grpc"
	"context"
	"errors"
	"log/slog"
)

type ApiStorage interface {
	GetUserApiSecret(ctx context.Context, userId int64, ApiPublic string) (string, error)
}
type CallApi interface {
	GetUserAccountInfo(ctx context.Context, PublicApiKey string, PrivateApiKey string) (*grpcFile.Response, error)
	GetFuturesTransactions(ctx context.Context, pageSize int64, PublicApiKey string, PrivateApiKey string) (*grpcFile.TransactionResponse, error)
}

type Bybit struct {
	apiStorage ApiStorage
	callApi    CallApi
	log        *slog.Logger
}

func NewBybit(apiStorage ApiStorage, api CallApi, log *slog.Logger) *Bybit {
	return &Bybit{apiStorage: apiStorage, callApi: api, log: log}
}

func (api *Bybit) SpotAccountInfo(ctx context.Context, publicApiKey string, userId int64) ([]*grpcFile.Coin, string, error) {
	//TODO validate parameters
	const op = "services.Bybit.SpotAccountInfo"
	log := api.log.With(slog.String("op", op),
		slog.Int64("userId", userId),
		slog.String("apiKey", publicApiKey))
	log.Info("getSpotAccountInfo")

	apiSecret, err := api.apiStorage.GetUserApiSecret(ctx, userId, publicApiKey)
	if err != nil {
		api.log.Error(err.Error())
		return nil, "", err
	}
	response, err := api.callApi.GetUserAccountInfo(ctx, publicApiKey, apiSecret)
	if err != nil {
		api.log.Error("callApi.GetUserAccountInfo", slog.String("err", err.Error()))
		return nil, "", err
	}
	if response.RetCode != 0 {
		api.log.Error("callApi.GetUserAccountInfo RetCode !=0", slog.String("err", response.RetMsg))
		return nil, "", errors.New(response.RetMsg)
	}
	return response.Result.List[0].Coins, response.Result.List[0].TotalWalletBalance, nil
}

func (api *Bybit) FuturesTransactions(ctx context.Context, pageSize int64, publicApiKey string, UserId int64) ([]*grpcFile.Transaction, error) {
	const op = "services.FuturesTransactions"
	log := api.log.With(slog.String("op", op),
		slog.Int64("UserId", UserId),
		slog.String("publicApiKey", publicApiKey))
	log.Info("getFuturesTransactions")
	apiSecret, err := api.apiStorage.GetUserApiSecret(ctx, UserId, publicApiKey)
	if err != nil {
		api.log.Error(err.Error())
		return nil, err
	}
	responce, err := api.callApi.GetFuturesTransactions(ctx, pageSize, publicApiKey, apiSecret)
	if err != nil {
		api.log.Error("callApi.GetUserAccountInfo", slog.String("err", err.Error()))
		return nil, err
	}

	return responce.Result.List, nil
}
