package callApi

import (
	grpcFile "bybit_api_servic_grpc/grpc"
	"bybit_api_servic_grpc/internal/callApi/models"
	"fmt"
	bybit "github.com/bybit-exchange/bybit.go.api"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/encoding/protojson"
	"log/slog"
)

type CallApi struct {
	log *slog.Logger
}

func (api *CallApi) GetUserAccountInfo(ctx context.Context, PublicApiKey string, PrivateApiKey string) (*grpcFile.Response, error) {
	const op = "CallApi.GetUserAccountInfo"
	client := bybit.NewBybitHttpClient(PublicApiKey, PrivateApiKey, bybit.WithBaseURL(bybit.MAINNET))
	params1 := map[string]interface{}{"accountType": "UNIFIED"}
	accountResults, err := client.NewUtaBybitServiceWithParams(params1).GetAccountWallet(context.Background())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	res := bybit.PrettyPrint(accountResults)
	result := parseJsonGetBalance(res)
	return &result, nil
}

func NewCallApi(log *slog.Logger) *CallApi {
	return &CallApi{log: log}
}

func (api *CallApi) GetFuturesTransactions(ctx context.Context, pageSize int, PublicApiKey string, PrivateApiKey string) (*models.Response, error) {
	//TODO implement function
	return nil, nil
}

func parseJsonGetBalance(res string) grpcFile.Response {
	var r grpcFile.Response
	if err := protojson.Unmarshal([]byte(res), &r); err != nil {
		fmt.Printf("Failed to parse JSON: %v\nOriginal JSON: %s\n", err, res)
	}
	return r
}
