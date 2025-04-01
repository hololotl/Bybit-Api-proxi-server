package models

import grpcFile "bybit_api_servic_grpc/grpc"

type Result struct {
	List []*grpcFile.SpotAccountResponse `json:"list"`
}

// Response представляет корневой объект JSON
type Response struct {
	RetCode    int      `json:"retCode"`
	RetMsg     string   `json:"retMsg"`
	Result     Result   `json:"result"`
	RetExtInfo struct{} `json:"retExtInfo"`
	Time       int64    `json:"time"`
}
