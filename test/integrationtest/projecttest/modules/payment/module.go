package payment

import (
	"flamingo.me/dingo"
	"github.com/lunarforge/flamingo_commerce/payment/interfaces"
)

type (
	// Module registers our fake payment profile
	Module struct{}
)

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMap((*interfaces.WebCartPaymentGateway)(nil), FakePaymentGateway).To(new(FakeGateway)).In(dingo.Singleton)
}
