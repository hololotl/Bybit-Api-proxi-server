package bybitApi

import (
	grpcFile "bybit_api_servic_grpc/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerApi struct {
	grpcFile.UnimplementedBybitServiceServer
	api ApiFunctions
}

type ApiFunctions interface {
	FuturesTransactions(ctx context.Context, pageSize int64, publicApiKey string, userId int64) ([]*grpcFile.Transaction, error)
	SpotAccountInfo(ctx context.Context, publicApiKey string, userId int64) ([]*grpcFile.Coin, string, error)
}

func RegisterApiFunctions(server *grpc.Server, apiFunctions ApiFunctions) {
	grpcFile.RegisterBybitServiceServer(server, &ServerApi{api: apiFunctions})
}

func (serverApi *ServerApi) FuturesTransactions(ctx context.Context, in *grpcFile.FuturesTransactionsRequest) (*grpcFile.TransactionResult, error) {
	if in.PublicApi == "" {
		return nil, status.Error(codes.InvalidArgument, "publicApi is required")
	}
	Transaction, err := serverApi.api.FuturesTransactions(ctx, in.PageSize, in.PublicApi, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcFile.TransactionResult{List: Transaction}, nil
}

func (serverApi *ServerApi) SpotAccountInfo(ctx context.Context, in *grpcFile.SpotAccountRequest) (*grpcFile.SpotAccountResponse, error) {
	if in.PublicApi == "" {
		return nil, status.Error(codes.InvalidArgument, "publicApi is required")
	}
	coins, TotalBalance, err := serverApi.api.SpotAccountInfo(ctx, in.PublicApi, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "SpotAccountInfo failed")
	}
	return &grpcFile.SpotAccountResponse{Coins: coins, TotalWalletBalance: TotalBalance}, nil
}
