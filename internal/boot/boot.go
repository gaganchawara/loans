package boot

import (
	"context"
	"github.com/gaganchawara/loans/pkg/errors"
	"github.com/gaganchawara/loans/pkg/logger"
)

func Initialize(ctx context.Context)  {
	logger.Get(ctx).Info("booting the application")

	// initializes error packages with Hooks
	errors.Initialize(logger.ErrorLogger())
}
