package cart_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/lunarforge/flamingo_commerce/cart"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(config.Map{
		"core.auth.web.debugController": false,
	}, new(cart.Module)); err != nil {
		t.Error(err)
	}
}
