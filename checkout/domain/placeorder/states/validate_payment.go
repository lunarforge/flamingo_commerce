package states

import (
	"context"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/payment/application"
	"go.opencensus.io/trace"
)

type (
	// ValidatePayment state
	ValidatePayment struct {
		paymentService *application.PaymentService
		validator      process.PaymentValidatorFunc
	}
)

var _ process.State = ValidatePayment{}

// Inject dependencies
func (v *ValidatePayment) Inject(
	paymentService *application.PaymentService,
	validator process.PaymentValidatorFunc,
) *ValidatePayment {
	v.paymentService = paymentService
	v.validator = validator

	return v
}

// Name get state name
func (ValidatePayment) Name() string {
	return "ValidatePayment"
}

// Run the state operations
func (v ValidatePayment) Run(ctx context.Context, p *process.Process) process.RunResult {
	ctx, span := trace.StartSpan(ctx, "placeorder/state/ValidatePayment/Run")
	defer span.End()

	return v.validator(ctx, p, v.paymentService)
}

// Rollback the state operations
func (v ValidatePayment) Rollback(ctx context.Context, _ process.RollbackData) error {
	return nil
}

// IsFinal if state is a final state
func (v ValidatePayment) IsFinal() bool {
	return false
}
