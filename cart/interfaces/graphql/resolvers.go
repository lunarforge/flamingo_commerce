package graphql

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"

	"github.com/lunarforge/flamingo_commerce/cart/application"
	"github.com/lunarforge/flamingo_commerce/cart/domain/cart"
	"github.com/lunarforge/flamingo_commerce/cart/domain/validation"
	"github.com/lunarforge/flamingo_commerce/cart/interfaces/graphql/dto"
	"github.com/lunarforge/flamingo_commerce/product/domain"
)

// CommerceCartQueryResolver resolver for carts
type CommerceCartQueryResolver struct {
	applicationCartReceiverService *application.CartReceiverService
	applicationCartService         *application.CartService
	restrictionService             *validation.RestrictionService
	productService                 domain.ProductService
}

// Inject dependencies
func (r *CommerceCartQueryResolver) Inject(
	applicationCartReceiverService *application.CartReceiverService,
	cartService *application.CartService,
	restrictionService *validation.RestrictionService,
	productService domain.ProductService,
) {
	r.applicationCartReceiverService = applicationCartReceiverService
	r.applicationCartService = cartService
	r.restrictionService = restrictionService
	r.productService = productService
}

// CommerceCart getter for queries
func (r *CommerceCartQueryResolver) CommerceCart(ctx context.Context) (*dto.DecoratedCart, error) {
	req := web.RequestFromContext(ctx)
	dc, err := r.applicationCartReceiverService.ViewDecoratedCart(ctx, req.Session())
	if err != nil {
		return nil, err
	}

	return dto.NewDecoratedCart(dc), nil
}

// CommerceCartValidator to trigger the cart validation service
func (r *CommerceCartQueryResolver) CommerceCartValidator(ctx context.Context) (*validation.Result, error) {
	session := web.SessionFromContext(ctx)

	decoratedCart, err := r.applicationCartReceiverService.ViewDecoratedCart(ctx, session)
	if err != nil {
		return nil, err
	}

	result := r.applicationCartService.ValidateCart(ctx, session, decoratedCart)

	return &result, nil
}

// CommerceCartQtyRestriction checks if given sku is restricted in terms of qty
func (r *CommerceCartQueryResolver) CommerceCartQtyRestriction(ctx context.Context, marketplaceCode string, variantCode *string, deliveryCode string) (*validation.RestrictionResult, error) {
	session := web.SessionFromContext(ctx)

	product, err := r.productService.Get(ctx, marketplaceCode)
	if err != nil {
		return nil, err
	}
	if variantCode != nil {
		if configurableProduct, ok := product.(domain.ConfigurableProduct); ok {
			product, err = configurableProduct.GetConfigurableWithActiveVariant(*variantCode)
			if err != nil {
				return nil, err
			}
		}
	}

	cart, err := r.applicationCartReceiverService.ViewCart(ctx, session)
	if err != nil {
		return nil, err
	}
	result := r.restrictionService.RestrictQty(ctx, session, product, cart, deliveryCode)
	return result, nil
}

// CartSplit returns graphql specific cart split
func (r *CommerceCartQueryResolver) CartSplit(_ context.Context, paymentSelection *cart.DefaultPaymentSelection) ([]*dto.PaymentSelectionSplit, error) {
	if paymentSelection == nil {
		return nil, nil
	}

	paymentSelectionSplit := make([]*dto.PaymentSelectionSplit, 0)
	for qualifier, charge := range paymentSelection.CartSplit() {
		paymentSelectionSplit = append(paymentSelectionSplit, &dto.PaymentSelectionSplit{
			Qualifier: qualifier,
			Charge:    charge,
		})
	}

	return paymentSelectionSplit, nil
}
