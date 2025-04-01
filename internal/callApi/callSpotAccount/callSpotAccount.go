package callSpotAccount

import (
	"bybit_api_servic_grpc/internal/callApi/models"
	"context"
	"encoding/json"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetUserAccountInfo(ctx context.Context, PublicApiKey string, PrivateApiKey string) (*models.Response, error) {
	client := bybit.NewBybitHttpClient(PublicApiKey, PrivateApiKey, bybit.WithBaseURL(bybit.MAINNET))
	params1 := map[string]interface{}{"accountType": "UNIFIED"}
	accountResults, err := client.NewUtaBybitServiceWithParams(params1).GetAccountWallet(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := bybit.PrettyPrint(accountResults)
	result := parseJsonGetBalance(res)
	return &result, nil
}

func parseJsonGetBalance(res string) models.Response {
	var r models.Response
	err := json.Unmarshal([]byte(res), &r)
	if err != nil {
		fmt.Println(err)
	}
	return r
}
