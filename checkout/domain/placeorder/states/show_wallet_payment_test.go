package states_test

import (
	"context"
	"testing"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"

	"github.com/stretchr/testify/assert"
)

func TestShowWalletPayment_IsFinal(t *testing.T) {
	s := states.ShowWalletPayment{}
	assert.False(t, s.IsFinal())
}

func TestShowWalletPayment_Name(t *testing.T) {
	s := states.ShowWalletPayment{}
	assert.Equal(t, "ShowWalletPayment", s.Name())
}

func TestShowWalletPayment_Rollback(t *testing.T) {
	s := states.ShowWalletPayment{}
	assert.Nil(t, s.Rollback(context.Background(), nil))
}

func TestShowWalletPayment_Run(t *testing.T) {
	s := states.ShowWalletPayment{}
	isCalled := false
	s.Inject(nil, func(_ context.Context, _ *process.Process, _ *application.PaymentService) process.RunResult {
		isCalled = true
		return process.RunResult{}
	})

	s.Run(context.Background(), nil)

	assert.True(t, isCalled)
}

func TestNewShowWalletPaymentStateData(t *testing.T) {
	assert.Equal(t,
		process.StateData(states.ShowWalletPaymentData{
			UsedPaymentMethod: "test",
		}),
		states.NewShowWalletPaymentStateData(states.ShowWalletPaymentData{
			UsedPaymentMethod: "test",
		}),
	)
}
