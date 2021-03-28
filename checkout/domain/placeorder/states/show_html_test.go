package states_test

import (
	"context"
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"

	"github.com/stretchr/testify/assert"
)

func TestShowHTML_IsFinal(t *testing.T) {
	s := states.ShowHTML{}
	assert.False(t, s.IsFinal())
}

func TestShowHTML_Name(t *testing.T) {
	s := states.ShowHTML{}
	assert.Equal(t, "ShowHTML", s.Name())
}

func TestShowHTML_Rollback(t *testing.T) {
	s := states.ShowHTML{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestShowHTML_Run(t *testing.T) {
	s := states.ShowHTML{}
	isCalled := false
	s.Inject(nil, func(_ context.Context, _ *process.Process, _ *application.PaymentService) process.RunResult {
		isCalled = true
		return process.RunResult{}
	})

	s.Run(context.Background(), nil)

	assert.True(t, isCalled)
}

func TestNewShowHTMLStateData(t *testing.T) {
	assert.Equal(t, process.StateData("<h2>test</h2>"), states.NewShowHTMLStateData("<h2>test</h2>"))
}
