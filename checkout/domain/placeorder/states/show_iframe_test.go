package states_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"

	"github.com/stretchr/testify/assert"
)

func TestShowIframe_IsFinal(t *testing.T) {
	s := states.ShowIframe{}
	assert.False(t, s.IsFinal())
}

func TestShowIframe_Name(t *testing.T) {
	s := states.ShowIframe{}
	assert.Equal(t, "ShowIframe", s.Name())
}

func TestShowIframe_Rollback(t *testing.T) {
	s := states.ShowIframe{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestShowIframe_Run(t *testing.T) {
	s := states.ShowIframe{}
	isCalled := false
	s.Inject(nil, func(_ context.Context, _ *process.Process, _ *application.PaymentService) process.RunResult {
		isCalled = true
		return process.RunResult{}
	})

	s.Run(context.Background(), nil)

	assert.True(t, isCalled)
}

func TestNewShowIframeStateData(t *testing.T) {
	assert.Equal(t, process.StateData(&url.URL{Host: "test.com"}), states.NewShowIframeStateData(&url.URL{Host: "test.com"}))
}
