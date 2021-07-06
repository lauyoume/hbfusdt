package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/lauyoume/hbfusdt/pkg/model/market"
)

func TestGetIndexs(t *testing.T) {
	ct := new(MarketClient)
	ct.Init("api.btcgateway.pro")

	indexs, err := ct.GetIndexs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ix := range indexs {
		fmt.Println(ix.ContractCode, ix.IndexPrice)
	}
}

func TestGetIndexOne(t *testing.T) {
	ct := new(MarketClient)
	ct.Init("api.btcgateway.pro")

	{
		ix, err := ct.GetIndex("BTC-USDT")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(ix.ContractCode, ix.IndexPrice.Round(2))
	}
	{
		ix, err := ct.GetIndex("DOT-USDT")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(ix.ContractCode, ix.IndexPrice)
	}
}

func TestGetKlineOne(t *testing.T) {
	ct := new(MarketClient)
	ct.Init("api.btcgateway.pro")

	{
		req := market.GetCandlestickOptionalRequest{
			Period: market.MIN1,
			Size:   100,
		}
		ix, err := ct.GetCandlestick("BTC-USDT", req)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, t := range ix {
			start := time.Unix(t.Id, 0)
			fmt.Println(start.Format("15:04:05"), t.Id, t.Open, t.Close, t.Amount, t.Vol, t.Count)
		}
	}
}
