package states_test

import (
	"context"
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"

	"github.com/stretchr/testify/assert"
)

func TestValidatePayment_IsFinal(t *testing.T) {
	assert.False(t, states.ValidatePayment{}.IsFinal())
}

func TestValidatePayment_Name(t *testing.T) {
	assert.Equal(t, "ValidatePayment", states.ValidatePayment{}.Name())
}

func TestValidatePayment_Rollback(t *testing.T) {
	s := states.ValidatePayment{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestValidatePayment_Run(t *testing.T) {
	s := states.ValidatePayment{}
	isCalled := false
	s.Inject(nil, func(_ context.Context, _ *process.Process, _ *application.PaymentService) process.RunResult {
		isCalled = true
		return process.RunResult{}
	})

	s.Run(context.Background(), nil)

	assert.True(t, isCalled)
}
