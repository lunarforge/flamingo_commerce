package placeorder

import (
	"net/url"

	cartDomain "github.com/lunarforge/flamingo_commerce/cart/domain/cart"
)

type (
	// StartPlaceOrderCommand triggers new place order
	StartPlaceOrderCommand struct {
		Cart      cartDomain.Cart
		ReturnURL *url.URL
	}

	// RefreshPlaceOrderCommand proceeds in place order process
	RefreshPlaceOrderCommand struct {
	}

	// CancelPlaceOrderCommand cancels current running process
	CancelPlaceOrderCommand struct {
	}
)
