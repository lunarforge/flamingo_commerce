package infrastructure

import (
	"context"

	customerDomain "github.com/lunarforge/flamingo_commerce/customer/domain"
	"flamingo.me/flamingo/v3/core/auth"
)

type (
	// NilCustomerServiceAdapter for CustomerService and CustomerIdentityService that returns always NotFound
	NilCustomerServiceAdapter struct{}
)

var _ customerDomain.CustomerIdentityService = new(NilCustomerServiceAdapter)

// GetByIdentity retrieves the authenticated customer by Identity
func (n *NilCustomerServiceAdapter) GetByIdentity(context.Context, auth.Identity) (customerDomain.Customer, error) {
	return nil, customerDomain.ErrCustomerNotFoundError
}
