package apiBybit

import (
	grpcFile "bybit_api_servic_grpc/grpc"
	"bybit_api_servic_grpc/internal/callApi/models"
	"context"
	"errors"
	"fmt"
)

type ApiStorage interface {
	GetUserApiSecret(ctx context.Context, userId int64, ApiPublic string) (string, error)
}
type CallApi interface {
	GetUserAccountInfo(ctx context.Context, PublicApiKey string, PrivateApiKey string) (*models.Response, error)
	GetFuturesTransactions(ctx context.Context, pageSize int, PublicApiKey string, PrivateApiKey string) (*models.Response, error)
}

type Bybit struct {
	apiStorage ApiStorage
	callApi    CallApi
}

func NewBybit(apiStorage ApiStorage, api CallApi) *Bybit {
	return &Bybit{apiStorage: apiStorage, callApi: api}
}

func (api *Bybit) SpotAccountInfo(ctx context.Context, publicApiKey string, userId int64) ([]*grpcFile.Coin, string, error) {
	//TODO validate parameters
	apiSecret, err := api.apiStorage.GetUserApiSecret(ctx, userId, publicApiKey)
	if err != nil {
		fmt.Println("error with getUserApiSecret", err)
	}
	response, err := api.callApi.GetUserAccountInfo(ctx, publicApiKey, apiSecret)
	if err != nil {
		fmt.Println("error with callSpotAccount.GetUserAccountInfo", err)
	}
	if response.RetCode != 0 {
		fmt.Println("RetCode =! 0, error with response")
		return nil, "", errors.New(response.RetMsg)
	}
	return response.Result.List[0].Coins, response.Result.List[0].TotalWalletBalance, nil
}

func (api *Bybit) FuturesTransactions(ctx context.Context, pageSize int, publicApiKey string) {

}
