package callApi

import (
	"bybit_api_servic_grpc/internal/callApi/models"
	"encoding/json"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
	"golang.org/x/net/context"
	"log/slog"
)

type CallApi struct {
	log *slog.Logger
}

func (api *CallApi) GetUserAccountInfo(ctx context.Context, PublicApiKey string, PrivateApiKey string) (*models.Response, error) {
	client := bybit.NewBybitHttpClient(PublicApiKey, PrivateApiKey, bybit.WithBaseURL(bybit.MAINNET))
	params1 := map[string]interface{}{"accountType": "UNIFIED"}
	accountResults, err := client.NewUtaBybitServiceWithParams(params1).GetAccountWallet(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := bybit.PrettyPrint(accountResults)
	fmt.Println(res)
	result := parseJsonGetBalance(res)
	fmt.Println(result)
	return &result, nil
}

func NewCallApi(log *slog.Logger) *CallApi {
	return &CallApi{log: log}
}

func (api *CallApi) GetFuturesTransactions(ctx context.Context, pageSize int, PublicApiKey string, PrivateApiKey string) (*models.Response, error) {
	//TODO implement function
	return nil, nil
}

func parseJsonGetBalance(res string) models.Response {
	var r models.Response
	if err := json.Unmarshal([]byte(res), &r); err != nil {
		fmt.Printf("Failed to parse JSON: %v\nOriginal JSON: %s\n", err, res)
	}
	return r
}
