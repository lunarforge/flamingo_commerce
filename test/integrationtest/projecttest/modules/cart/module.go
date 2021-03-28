package cart

import (
	"flamingo.me/dingo"
	"github.com/lunarforge/flamingo_commerce/cart"
	"github.com/lunarforge/flamingo_commerce/cart/domain/validation"
	"github.com/lunarforge/flamingo_commerce/cart/infrastructure"
)

type (
	// Module for integration testing
	Module struct{}
)

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Override((*infrastructure.VoucherHandler)(nil), "").To(&FakeVoucherHandler{})
	injector.BindMulti((*validation.MaxQuantityRestrictor)(nil)).To(FakeQtyRestrictor{})
	injector.Bind(new(validation.PaymentSelectionValidator)).To(new(FakePaymentSelectionValidator))
}

// Depends on other modules
func (m *Module) Depends() []dingo.Module {
	return []dingo.Module{
		&cart.Module{},
	}
}
