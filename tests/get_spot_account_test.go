package tests

import (
	grpcApi "bybit_api_servic_grpc/grpc"
	"bybit_api_servic_grpc/tests/suite"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_get_spot_account_test_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)
	userId := 2
	PublicApiKey := "3w3LwpWhSbr9wIzKHA"
	resp, err := st.AppClient.SpotAccountInfo(ctx, &grpcApi.SpotAccountRequest{UserId: int64(userId), PublicApi: PublicApiKey})
	if err != nil {
		fmt.Printf("SpotAccountInfo error: %v\n", err)
	}
	fmt.Println("resp:", resp.Coins)
	require.NoError(t, err)
	assert.NotEmpty(t, resp)
}
