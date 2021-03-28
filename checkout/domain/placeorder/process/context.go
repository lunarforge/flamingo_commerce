package process

import (
	"context"
	"net/url"

	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/checkout/application"
)

type (
	// Context contains information (state etc) about a place order process
	Context struct {
		UUID               string
		CurrentStateName   string
		CurrentStateData   StateData
		PlaceOrderInfo     *application.PlaceOrderInfo
		Cart               cart.Cart
		ReturnURL          *url.URL
		RollbackReferences []RollbackReference
		FailedReason       FailedReason
	}
	// StateData holding state relevant data
	StateData interface{}

	// ContextStore can persist process Context instances
	ContextStore interface {
		Store(ctx context.Context, key string, placeOrderContext Context) error
		Get(ctx context.Context, key string) (Context, bool)
		Delete(ctx context.Context, key string) error
	}
)
