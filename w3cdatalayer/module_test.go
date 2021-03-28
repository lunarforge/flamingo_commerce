package w3cdatalayer_test

import (
	"testing"

	"github.com/lunarforge/flamingo_commerce/w3cdatalayer"
	"flamingo.me/flamingo/v3/framework/config"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(nil, new(w3cdatalayer.Module)); err != nil {
		t.Error(err)
	}
}
