// +build integration

package restapi_test

import (
	"flag"
	"os"
	"testing"

	"github.com/lunarforge/flamingo_commerce/test/integrationtest/projecttest/helper"
)

var FlamingoURL string

// TestMain used for setup and teardown
func TestMain(m *testing.M) {
	flag.Parse()
	info := helper.BootupDemoProject("../../config/")
	FlamingoURL = info.BaseURL
	result := m.Run()
	info.ShutdownFunc()
	os.Exit(result)
}
