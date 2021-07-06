package client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/lauyoume/hbfusdt/internal"
	"github.com/lauyoume/hbfusdt/internal/requestbuilder"
	"github.com/lauyoume/hbfusdt/pkg/model"
	"github.com/lauyoume/hbfusdt/pkg/model/market"
)

// Responsible to get market information
type MarketClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

// Initializer
func (p *MarketClient) Init(host string) *MarketClient {
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

func (client *MarketClient) GetIndexs() ([]market.Index, error) {
	url := client.publicUrlBuilder.Build("/linear-swap-api/v1/swap_index", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Get Contract Index
func (client *MarketClient) GetIndex(symbol string) (*market.Index, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	url := client.publicUrlBuilder.Build("/linear-swap-api/v1/swap_index", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && len(result.Data) > 0 {
		index := result.Data[0]
		return &index, nil
	}
	return nil, errors.New(getResp)
}

// Retrieves all klines in a specific range.
func (client *MarketClient) GetCandlestick(symbol string, optionalRequest market.GetCandlestickOptionalRequest) ([]market.Candlestick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)
	if optionalRequest.Period != "" {
		request.AddParam("period", optionalRequest.Period)
	}
	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	} else if optionalRequest.From != 0 && optionalRequest.To != 0 {
		request.AddParam("from", strconv.Itoa(optionalRequest.From))
		request.AddParam("to", strconv.Itoa(optionalRequest.To))
	}

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/history/kline", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetCandlestickResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Retrieves the latest ticker with some important 24h aggregated market data.
func (client *MarketClient) GetLast24hCandlestickAskBid(symbol string) (*market.CandlestickAskBid, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("symbol", symbol)

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/detail/merged", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLast24hCandlestickAskBidResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}

// Retrieve the latest tickers for all supported pairs.
func (client *MarketClient) GetAllSymbolLast24hCandlesticksAskBid() ([]market.SymbolCandlestick, error) {

	request := new(model.GetRequest).Init()

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/detail/batch_merged", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetAllSymbolsLast24hCandlesticksAskBidResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Retrieves the current order book of a specific pair
func (client *MarketClient) GetDepth(symbol string, step string) (*market.Depth, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)
	request.AddParam("type", step)

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/depth", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetDepthResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}

// Retrieves the latest trade with its price, volume, and direction.
func (client *MarketClient) GetLatestTrade(symbol string) (*market.TradeTick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLatestTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}

// Retrieves the most recent trades with their price, volume, and direction.
func (client *MarketClient) GetHistoricalTrade(symbol string, optionalRequest market.GetHistoricalTradeOptionalRequest) ([]market.TradeTick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)
	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	}

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/history/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetHistoricalTradeResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Retrieves the summary of trading in the market for the last 24 hours.
func (client *MarketClient) GetLast24hCandlestick(symbol string) (*market.Candlestick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/detail/merged", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLast24hCandlestick{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}
