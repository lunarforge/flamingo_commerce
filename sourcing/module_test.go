package sourcing_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/lunarforge/flamingo_commerce/sourcing"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(config.Map{
		"core.auth.web.debugController": false,
	}, new(sourcing.Module)); err != nil {
		t.Error(err)
	}
}
