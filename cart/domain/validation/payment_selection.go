package validation

import (
	"context"

	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/cart/domain/decorator"
)

type (
	// PaymentSelectionValidator decides if the PaymentSelection is valid
	PaymentSelectionValidator interface {
		Validate(ctx context.Context, cart *decorator.DecoratedCart, selection cart.PaymentSelection) error
	}
)
