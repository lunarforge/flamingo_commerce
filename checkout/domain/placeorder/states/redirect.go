package states

import (
	"context"
	"encoding/gob"
	"net/url"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"github.com/lunarforge/flamingo_commerce/payment/application"
	"go.opencensus.io/trace"
)

type (
	// Redirect state
	Redirect struct {
		paymentService *application.PaymentService
		validator      process.PaymentValidatorFunc
	}
)

var _ process.State = Redirect{}

func init() {
	gob.Register(&url.URL{})
}

// NewRedirectStateData creates data required for this state
func NewRedirectStateData(url *url.URL) process.StateData {
	return process.StateData(url)
}

// Inject dependencies
func (r *Redirect) Inject(
	paymentService *application.PaymentService,
	validator process.PaymentValidatorFunc,
) *Redirect {
	r.paymentService = paymentService
	r.validator = validator

	return r
}

// Name get state name
func (Redirect) Name() string {
	return "Redirect"
}

// Run the state operations
func (r Redirect) Run(ctx context.Context, p *process.Process) process.RunResult {
	ctx, span := trace.StartSpan(ctx, "placeorder/state/Redirect/Run")
	defer span.End()

	return r.validator(ctx, p, r.paymentService)
}

// Rollback the state operations
func (r Redirect) Rollback(ctx context.Context, _ process.RollbackData) error {
	return nil
}

// IsFinal if state is a final state
func (r Redirect) IsFinal() bool {
	return false
}
