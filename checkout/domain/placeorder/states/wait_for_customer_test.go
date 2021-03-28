package states_test

import (
	"context"
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"
	"github.com/stretchr/testify/assert"
)

func TestWaitForCustomer_IsFinal(t *testing.T) {
	s := states.WaitForCustomer{}
	assert.False(t, s.IsFinal())
}

func TestWaitForCustomer_Name(t *testing.T) {
	s := states.WaitForCustomer{}
	assert.Equal(t, "WaitForCustomer", s.Name())
}

func TestWaitForCustomer_Rollback(t *testing.T) {
	s := states.WaitForCustomer{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestWaitForCustomer_Run(t *testing.T) {
	s := states.WaitForCustomer{}
	isCalled := false
	s.Inject(nil, func(_ context.Context, _ *process.Process, _ *application.PaymentService) process.RunResult {
		isCalled = true
		return process.RunResult{}
	})

	s.Run(context.Background(), nil)

	assert.True(t, isCalled)
}
