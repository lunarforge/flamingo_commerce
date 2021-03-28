package checkout_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/lunarforge/flamingo_commerce/checkout"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(config.Map{
		"core.auth.web.debugController": false,
	}, new(checkout.Module)); err != nil {
		t.Error(err)
	}
}
