package sourcing

import (
	"flamingo.me/dingo"
	"github.com/lunarforge/flamingo_commerce/cart/domain/validation"
	restrictors "github.com/lunarforge/flamingo_commerce/sourcing/domain/restrictor"

	"github.com/lunarforge/flamingo_commerce/cart"
	"github.com/lunarforge/flamingo_commerce/sourcing/application"
	"github.com/lunarforge/flamingo_commerce/sourcing/domain"
)

type (
	// Module registers sourcing module
	Module struct {
		useDefaultSourcingService bool
		enableQtyRestrictor       bool
	}
)

// Inject dependencies
func (m *Module) Inject(
	config *struct {
		UseDefaultSourcingService bool `inject:"config:commerce.sourcing.useDefaultSourcingService,optional"`
		EnableQtyRestrictor       bool `inject:"config:commerce.sourcing.enableQtyRestrictor,optional"`
	},
) {

	if config != nil {
		m.useDefaultSourcingService = config.UseDefaultSourcingService
		m.enableQtyRestrictor = config.EnableQtyRestrictor
	}

}

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	if m.useDefaultSourcingService {
		injector.Bind(new(domain.SourcingService)).To(domain.DefaultSourcingService{})
	}

	if m.enableQtyRestrictor {
		injector.Bind(new(validation.MaxQuantityRestrictor)).To(restrictors.Restrictor{})
	}

	injector.Bind(new(application.SourcingApplication)).To(application.Service{})
}

// Depends on other modules
func (m *Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(cart.Module),
	}
}

// CueConfig defines the sourcing module configuration
func (m *Module) CueConfig() string {
	return `
commerce: {
	sourcing: {
		useDefaultSourcingService: bool | *true
		enableQtyRestrictor: bool | *false
	}
}
`
}
