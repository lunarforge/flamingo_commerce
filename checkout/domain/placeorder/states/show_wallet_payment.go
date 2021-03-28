package states

import (
	"context"
	"encoding/gob"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/payment/application"
	"github.com/lunarforge/flamingo_commerce/payment/domain"

	"go.opencensus.io/trace"
)

type (
	// ShowWalletPayment state
	ShowWalletPayment struct {
		paymentService *application.PaymentService
		validator      process.PaymentValidatorFunc
	}

	// ShowWalletPaymentData holds details regarding the wallet payment
	ShowWalletPaymentData domain.WalletDetails
)

func init() {
	gob.Register(ShowWalletPaymentData{})
}

var _ process.State = ShowWalletPayment{}

// NewShowWalletPaymentStateData creates new StateData with (persisted) Data required for this state
func NewShowWalletPaymentStateData(walletDetails ShowWalletPaymentData) process.StateData {
	return process.StateData(walletDetails)
}

// Inject dependencies
func (pr *ShowWalletPayment) Inject(
	paymentService *application.PaymentService,
	validator process.PaymentValidatorFunc,
) *ShowWalletPayment {
	pr.paymentService = paymentService
	pr.validator = validator

	return pr
}

// Name get state name
func (ShowWalletPayment) Name() string {
	return "ShowWalletPayment"
}

// Run the state operations
func (pr ShowWalletPayment) Run(ctx context.Context, p *process.Process) process.RunResult {
	ctx, span := trace.StartSpan(ctx, "placeorder/state/ShowWalletPayment/Run")
	defer span.End()

	return pr.validator(ctx, p, pr.paymentService)
}

// Rollback the state operations
func (pr ShowWalletPayment) Rollback(ctx context.Context, _ process.RollbackData) error {
	return nil
}

// IsFinal if state is a final state
func (pr ShowWalletPayment) IsFinal() bool {
	return false
}
