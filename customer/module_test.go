package customer_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/lunarforge/flamingo_commerce/customer"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(config.Map{
		"commerce.customer.useNilCustomerAdapter": true,
		"core.auth.web.debugController":           false,
	}, new(customer.Module)); err != nil {
		t.Error(err)
	}
}
