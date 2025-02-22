package states

import (
	"context"

	"github.com/lunarforge/flamingo_commerce/checkout/domain/placeorder/process"
	"go.opencensus.io/trace"
)

type (
	// Failed state
	Failed struct {
		Reason process.FailedReason
	}
)

var _ process.State = Failed{}

// Name get state name
func (f Failed) Name() string {
	return "Failed"
}

// Run the state operations
func (f Failed) Run(ctx context.Context, _ *process.Process) process.RunResult {
	_, span := trace.StartSpan(ctx, "placeorder/state/Failed/Run")
	defer span.End()

	return process.RunResult{}
}

// Rollback the state operations
func (f Failed) Rollback(ctx context.Context, _ process.RollbackData) error {
	return nil
}

// IsFinal if state is a final state
func (f Failed) IsFinal() bool {
	return true
}
