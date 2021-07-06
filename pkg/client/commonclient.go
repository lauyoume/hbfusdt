package client

import (
	"encoding/json"
	"errors"

	"github.com/lauyoume/hbfusdt/internal"
	"github.com/lauyoume/hbfusdt/internal/requestbuilder"
	"github.com/lauyoume/hbfusdt/pkg/model"
	"github.com/lauyoume/hbfusdt/pkg/model/common"
)

// Responsible to get common information
type CommonClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

// Initializer
func (p *CommonClient) Init(host string) *CommonClient {
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

func (p *CommonClient) GetSystemStatus() (string, error) {
	url := "https://status.huobigroup.com/api/v2/summary.json"
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return "", getErr
	}

	return getResp, nil
}

// Returns current market status
func (p *CommonClient) GetMarketStatus() (*common.MarketStatus, error) {
	url := p.publicUrlBuilder.Build("/heartbeat", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetMarketStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && &result.Data != nil {
		return &result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Get all Supported Trading Symbol
// This endpoint returns all Huobi's supported trading symbol.
func (p *CommonClient) GetSymbols() ([]common.Symbol, error) {
	url := p.publicUrlBuilder.Build("/linear-swap-api/v1/swap_contract_info", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetSymbolsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Get all Supported Trading Symbol
// This endpoint returns all Huobi's supported trading symbol.
func (p *CommonClient) GetSymbol(symbol string) (*common.Symbol, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	url := p.publicUrlBuilder.Build("/linear-swap-api/v1/swap_contract_info", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := common.GetSymbolsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && len(result.Data) > 0 {
		ret := result.Data[0]
		return &ret, nil
	}
	return nil, errors.New(getResp)
}

// Get Current Timestamp
// This endpoint returns the current timestamp, i.e. the number of milliseconds that have elapsed since 00:00:00 UTC on 1 January 1970.
func (p *CommonClient) GetTimestamp() (int64, error) {
	url := p.publicUrlBuilder.Build("/api/v1/timestamp", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return 0, getErr
	}

	result := common.GetTimestampResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)

	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Status == "ok" && result.Ts != 0 {
		return result.Ts, nil
	}
	return 0, errors.New(getResp)
}
