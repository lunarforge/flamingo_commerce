package states_test

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cartDomain "github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/cart/domain/placeorder"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/states"
	"github.com/lunarforge/flamingo_commerce/payment/application"
	"github.com/lunarforge/flamingo_commerce/payment/domain"
	"github.com/lunarforge/flamingo_commerce/payment/interfaces"
	"github.com/lunarforge/flamingo_commerce/payment/interfaces/mocks"
	price "github.com/lunarforge/flamingo_commerce/price/domain"
)

func TestCreatePayment_Run(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		factory := provideProcessFactory(t)

		cart := provideCartWithPaymentSelection(t)
		p, _ := factory.New(&url.URL{}, cart)

		state := states.CreatePayment{}

		expectedPayment := &placeorder.Payment{Gateway: "test"}
		gateway := &mocks.WebCartPaymentGateway{}
		gateway.On("StartFlow", mock.Anything, mock.Anything, p.Context().UUID, p.Context().ReturnURL).Return(&domain.FlowResult{EarlyPlaceOrder: true}, nil).Once()
		gateway.On("OrderPaymentFromFlow", mock.Anything, mock.Anything, p.Context().UUID).Return(expectedPayment, nil).Once()
		paymentService := paymentServiceHelper(t, gateway)

		state.Inject(paymentService)

		expectedResult := process.RunResult{
			RollbackData: states.CreatePaymentRollbackData{Gateway: expectedPayment.Gateway, PaymentID: expectedPayment.PaymentID},
		}
		result := state.Run(context.Background(), p)
		assert.Equal(t, result, expectedResult)
		assert.Equal(t, p.Context().CurrentStateName, states.CompleteCart{}.Name())
		gateway.AssertExpectations(t)
	})

	t.Run("error missing payment selection", func(t *testing.T) {
		factory := provideProcessFactory(t)

		cart := cartDomain.Cart{}

		p, _ := factory.New(&url.URL{}, cart)

		state := states.CreatePayment{}

		paymentService := paymentServiceHelper(t, nil)

		state.Inject(paymentService)

		result := state.Run(context.Background(), p)
		assert.NotNil(t, result.Failed, "Missing PaymentSelection in cart should lead to an error")
	})

	t.Run("error during gateway.StartFlow", func(t *testing.T) {
		factory := provideProcessFactory(t)

		cart := provideCartWithPaymentSelection(t)

		p, _ := factory.New(&url.URL{}, cart)

		state := states.CreatePayment{}

		expectedError := errors.New("StartFlow payment error")

		gateway := &mocks.WebCartPaymentGateway{}
		gateway.On("StartFlow", mock.Anything, mock.Anything, p.Context().UUID, p.Context().ReturnURL).Return(nil, expectedError).Once()
		paymentService := paymentServiceHelper(t, gateway)
		state.Inject(paymentService)

		expectedResult := process.RunResult{
			Failed: process.PaymentErrorOccurredReason{Error: expectedError.Error()},
		}
		assert.Equal(t, state.Run(context.Background(), p), expectedResult)
		gateway.AssertExpectations(t)
	})

	t.Run("error during gateway.OrderPaymentFromFlow", func(t *testing.T) {
		factory := provideProcessFactory(t)

		cart := provideCartWithPaymentSelection(t)

		p, _ := factory.New(&url.URL{}, cart)

		state := states.CreatePayment{}

		expectedError := errors.New("OrderPaymentFromFlow payment error")

		gateway := &mocks.WebCartPaymentGateway{}
		gateway.On("StartFlow", mock.Anything, mock.Anything, p.Context().UUID, p.Context().ReturnURL).Return(&domain.FlowResult{}, nil).Once()
		gateway.On("OrderPaymentFromFlow", mock.Anything, mock.Anything, p.Context().UUID).Return(nil, expectedError).Once()

		paymentService := paymentServiceHelper(t, gateway)
		state.Inject(paymentService)

		expectedResult := process.RunResult{
			Failed: process.PaymentErrorOccurredReason{Error: expectedError.Error()},
		}
		assert.Equal(t, state.Run(context.Background(), p), expectedResult)
		gateway.AssertExpectations(t)
	})
}

func provideProcessFactory(t *testing.T) *process.Factory {
	t.Helper()
	factory := &process.Factory{}
	factory.Inject(
		func() *process.Process {
			return &process.Process{}
		},
		&struct {
			StartState  process.State `inject:"startState"`
			FailedState process.State `inject:"failedState"`
		}{
			StartState: &states.New{},
		},
	)
	return factory
}

func provideCartWithPaymentSelection(t *testing.T) cartDomain.Cart {
	t.Helper()
	cart := cartDomain.Cart{}
	paymentSelection, err := cartDomain.NewDefaultPaymentSelection("test", map[string]string{price.ChargeTypeMain: "main"}, cart)
	require.NoError(t, err)
	cart.PaymentSelection = paymentSelection
	return cart
}

func paymentServiceHelper(t *testing.T, gateway interfaces.WebCartPaymentGateway) *application.PaymentService {
	t.Helper()
	paymentService := &application.PaymentService{}

	paymentService.Inject(func() map[string]interfaces.WebCartPaymentGateway {
		return map[string]interfaces.WebCartPaymentGateway{
			"test": gateway,
		}
	})
	return paymentService
}

func TestCreatePayment_IsFinal(t *testing.T) {
	state := states.CreatePayment{}
	assert.False(t, state.IsFinal())
}

func TestCreatePayment_Rollback(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		state := states.CreatePayment{}

		var data interface{}

		payment := &placeorder.Payment{
			Gateway:            "test",
			Transactions:       nil,
			RawTransactionData: nil,
			PaymentID:          "1234",
		}

		data = states.CreatePaymentRollbackData{Gateway: payment.Gateway, PaymentID: payment.PaymentID}

		gateway := &mocks.WebCartPaymentGateway{}
		gateway.On("CancelOrderPayment", mock.Anything, payment).Return(nil).Once()
		paymentService := paymentServiceHelper(t, gateway)
		state.Inject(paymentService)

		result := state.Rollback(context.Background(), data)
		assert.Nil(t, result)
		gateway.AssertExpectations(t)
	})

	t.Run("RollbackData not of type", func(t *testing.T) {
		state := states.CreatePayment{}

		assert.Error(t, state.Rollback(context.Background(), "string"))
	})

	t.Run("Error during payment selection", func(t *testing.T) {
		state := states.CreatePayment{}

		var data interface{}

		payment := &placeorder.Payment{
			Gateway: "non-existing",
		}

		data = states.CreatePaymentRollbackData{Gateway: payment.Gateway, PaymentID: payment.PaymentID}

		paymentService := paymentServiceHelper(t, nil)
		state.Inject(paymentService)
		assert.Error(t, state.Rollback(context.Background(), data), "Missing payment selection / gateway should lead to an error")
	})

	t.Run("Error during CancelOrderPayment", func(t *testing.T) {
		state := states.CreatePayment{}

		var data interface{}

		payment := &placeorder.Payment{
			Gateway:            "test",
			Transactions:       nil,
			RawTransactionData: nil,
			PaymentID:          "1234",
		}

		data = states.CreatePaymentRollbackData{Gateway: payment.Gateway, PaymentID: payment.PaymentID}

		gateway := &mocks.WebCartPaymentGateway{}
		expectedError := errors.New("generic payment error")
		gateway.On("CancelOrderPayment", mock.Anything, payment).Return(expectedError).Once()
		paymentService := paymentServiceHelper(t, gateway)
		state.Inject(paymentService)
		assert.EqualError(t, state.Rollback(context.Background(), data), expectedError.Error())
		gateway.AssertExpectations(t)
	})

}
