package suite

import (
	grpcApi "bybit_api_servic_grpc/grpc"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strconv"
	"testing"
	"time"
)

type Suite struct {
	*testing.T
	AppClient grpcApi.BybitServiceClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()
	timeout := time.Duration(1 * time.Minute)
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})
	grpcHost := "localhost"
	grpcAddress := net.JoinHostPort(grpcHost, strconv.Itoa(8083))

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress,
		// Используем insecure-коннект для тестов
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("пизда бля горим")
		t.Fatalf("grpc server connection failed: %v", err)
	}
	authClient := grpcApi.NewBybitServiceClient(cc)
	return ctx, &Suite{t, authClient}
}
