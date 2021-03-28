// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cart "github.com/lunarforge/flamingo_commerce/cart/domain/cart"
import context "context"
import domain "github.com/lunarforge/flamingo_commerce/payment/domain"

import mock "github.com/stretchr/testify/mock"
import placeorder "github.com/lunarforge/flamingo_commerce/cart/domain/placeorder"
import url "net/url"

// WebCartPaymentGateway is an autogenerated mock type for the WebCartPaymentGateway type
type WebCartPaymentGateway struct {
	mock.Mock
}

// CancelOrderPayment provides a mock function with given fields: ctx, cartPayment
func (_m *WebCartPaymentGateway) CancelOrderPayment(ctx context.Context, cartPayment *placeorder.Payment) error {
	ret := _m.Called(ctx, cartPayment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *placeorder.Payment) error); ok {
		r0 = rf(ctx, cartPayment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConfirmResult provides a mock function with given fields: ctx, _a1, cartPayment
func (_m *WebCartPaymentGateway) ConfirmResult(ctx context.Context, _a1 *cart.Cart, cartPayment *placeorder.Payment) error {
	ret := _m.Called(ctx, _a1, cartPayment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *placeorder.Payment) error); ok {
		r0 = rf(ctx, _a1, cartPayment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlowStatus provides a mock function with given fields: ctx, _a1, correlationID
func (_m *WebCartPaymentGateway) FlowStatus(ctx context.Context, _a1 *cart.Cart, correlationID string) (*domain.FlowStatus, error) {
	ret := _m.Called(ctx, _a1, correlationID)

	var r0 *domain.FlowStatus
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *domain.FlowStatus); ok {
		r0 = rf(ctx, _a1, correlationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FlowStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) error); ok {
		r1 = rf(ctx, _a1, correlationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Methods provides a mock function with given fields:
func (_m *WebCartPaymentGateway) Methods() []domain.Method {
	ret := _m.Called()

	var r0 []domain.Method
	if rf, ok := ret.Get(0).(func() []domain.Method); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Method)
		}
	}

	return r0
}

// OrderPaymentFromFlow provides a mock function with given fields: ctx, _a1, correlationID
func (_m *WebCartPaymentGateway) OrderPaymentFromFlow(ctx context.Context, _a1 *cart.Cart, correlationID string) (*placeorder.Payment, error) {
	ret := _m.Called(ctx, _a1, correlationID)

	var r0 *placeorder.Payment
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *placeorder.Payment); ok {
		r0 = rf(ctx, _a1, correlationID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*placeorder.Payment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) error); ok {
		r1 = rf(ctx, _a1, correlationID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartFlow provides a mock function with given fields: ctx, _a1, correlationID, returnURL
func (_m *WebCartPaymentGateway) StartFlow(ctx context.Context, _a1 *cart.Cart, correlationID string, returnURL *url.URL) (*domain.FlowResult, error) {
	ret := _m.Called(ctx, _a1, correlationID, returnURL)

	var r0 *domain.FlowResult
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, *url.URL) *domain.FlowResult); ok {
		r0 = rf(ctx, _a1, correlationID, returnURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FlowResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, *url.URL) error); ok {
		r1 = rf(ctx, _a1, correlationID, returnURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
