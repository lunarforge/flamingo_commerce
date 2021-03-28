package dto

import (
	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/price/domain"
)

// PaymentSelectionSplit is a GraphQL specific representation of `cart.PaymentSplit`
type PaymentSelectionSplit struct {
	Qualifier cart.SplitQualifier
	Charge    domain.Charge
}
