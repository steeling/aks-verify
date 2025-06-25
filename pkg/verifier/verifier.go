package verifier

import (
	"context"

	"github.com/steeling/aks-verify/pkg/result"
)

type Verifier interface {
	Name() string
	Description() string
	Run(ctx context.Context) result.Report
}
