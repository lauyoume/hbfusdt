package client

import (
	"fmt"
	"testing"
)

func TestGetTimeStamp(t *testing.T) {
	ct := new(CommonClient)
	ct.Init("api.btcgateway.pro")

	fmt.Println(ct.GetTimestamp())
}

func TestGetSymbols(t *testing.T) {
	ct := new(CommonClient)
	ct.Init("api.btcgateway.pro")

	symbols, err := ct.GetSymbols()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range symbols {
		fmt.Println(s.Symbol, s.ContractCode, s.PriceTick)
	}
}

func TestGetSymbol(t *testing.T) {
	ct := new(CommonClient)
	ct.Init("api.btcgateway.pro")

	s, err := ct.GetSymbol("BTC-USDT")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s.Symbol, s.ContractCode, s.PriceTick, s.ContractSize)
}
