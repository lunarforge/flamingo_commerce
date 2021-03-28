package dto

import (
	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
)

type (
	//Taxes – provides custom graphql interface methods
	Taxes struct {
		Items []cart.Tax
	}
)

// GetByType - returns tax by given type
func (t Taxes) GetByType(taxType string) cart.Tax {
	for _, tax := range t.Items {
		if tax.Type == taxType {
			return tax
		}
	}
	return cart.Tax{}
}
