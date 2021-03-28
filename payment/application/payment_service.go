package application

import (
	"errors"

	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/payment/interfaces"
)

type (
	// PaymentService defines the payment service
	PaymentService struct {
		webCartPaymentGateways map[string]interfaces.WebCartPaymentGateway
	}
)

// Inject dependencies
func (ps *PaymentService) Inject(
	webCartPaymentGatewayProvider interfaces.WebCartPaymentGatewayProvider,
) {
	ps.webCartPaymentGateways = webCartPaymentGatewayProvider()
}

// PaymentGateway tries to get the supplied payment gateway by code from the registered payment gateways
func (ps *PaymentService) PaymentGateway(paymentGatewayCode string) (interfaces.WebCartPaymentGateway, error) {
	gateway, ok := ps.webCartPaymentGateways[paymentGatewayCode]
	if !ok {
		return nil, errors.New("Payment gateway " + paymentGatewayCode + " not found")
	}

	return gateway, nil
}

// AvailablePaymentGateways returns the list of registered WebCartPaymentGateway
func (ps *PaymentService) AvailablePaymentGateways() map[string]interfaces.WebCartPaymentGateway {
	return ps.webCartPaymentGateways
}

// PaymentGatewayByCart tries to get the payment gateway from the supllied cart
func (ps *PaymentService) PaymentGatewayByCart(cart cart.Cart) (interfaces.WebCartPaymentGateway, error) {
	if cart.PaymentSelection == nil {
		return nil, errors.New("PaymentSelection not set")
	}

	return ps.PaymentGateway(cart.PaymentSelection.Gateway())
}
