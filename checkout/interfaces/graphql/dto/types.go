package dto

import (
	"github.com/lunarforge/flamingo_commerce/cart/domain/placeorder"
	"github.com/lunarforge/flamingo_commerce/cart/interfaces/graphql/dto"
	"github.com/lunarforge/flamingo_commerce/checkout/application"
)

type (
	// StartPlaceOrderResult result of start place order
	StartPlaceOrderResult struct {
		UUID string
	}

	// PlaceOrderContext infos
	PlaceOrderContext struct {
		Cart       *dto.DecoratedCart
		OrderInfos *PlacedOrderInfos
		State      State
		UUID       string
	}

	// PlacedOrderInfos infos
	PlacedOrderInfos struct {
		PaymentInfos        []application.PlaceOrderPaymentInfo
		PlacedOrderInfos    []placeorder.PlacedOrderInfo
		Email               string
		PlacedDecoratedCart *dto.DecoratedCart
	}
)
