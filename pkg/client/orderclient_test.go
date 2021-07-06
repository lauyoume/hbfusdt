package client

import (
	"fmt"
	"testing"

	"github.com/lauyoume/hbfusdt/pkg/model/order"
)

func TestPlaceOrder(t *testing.T) {

	key := "87a7624d-6d059250-f7b95c81-dab4c45e6f"
	secret := "c6c0aae8-09aa026c-4c5a6660-448af"
	ct := new(OrderClient)
	ct.Init(key, secret, "api.btcgateway.pro")

	// 开多
	req := order.PlaceOrderRequest{
		Symbol:         "BTC-USDT",
		Direction:      order.BUY,
		Offset:         order.OPEN,
		OrderPriceType: order.TYPE_OPPONENT,
		LeverRate:      75,
		Volume:         1,
	}
	ret, err := ct.PlaceOrder(&req)
	fmt.Println(ret, err)
}

func TestPlaceOrderClose(t *testing.T) {

	key := "87a7624d-6d059250-f7b95c81-dab4c45e6f"
	secret := "c6c0aae8-09aa026c-4c5a6660-448af"
	ct := new(OrderClient)
	ct.Init(key, secret, "api.btcgateway.pro")

	// 平多
	req := order.PlaceOrderRequest{
		Symbol:         "BTC-USDT",
		Direction:      order.SELL,
		Offset:         order.CLOSE,
		OrderPriceType: order.TYPE_OPPONENT,
		LeverRate:      75,
		Volume:         1,
	}
	ret, err := ct.PlaceOrder(&req)
	fmt.Println(ret, err)
}
