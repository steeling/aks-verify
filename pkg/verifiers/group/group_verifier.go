package group

import (
	"context"

	"github.com/steeling/aks-verify/pkg/result"
	"github.com/steeling/aks-verify/pkg/verifier"
)

type Verifier struct {
	Name        string
	Description string
	Verifiers   []verifier.Verifier
}

func (v *Verifier) Run(ctx context.Context) result.Report {
	subreports := make([]result.Report, len(v.Verifiers))
	status := result.StatusSuccess
	for i, subVerifier := range v.Verifiers {
		report := subVerifier.Run(ctx)
		subreports[i] = report
		if report.Status == result.StatusFailure {
			status = result.StatusFailure
		}
		if report.Status == result.StatusError {
			status = result.StatusError
		}
	}
	return result.Report{
		Name:        v.Name,
		Description: v.Description,
		SubReports:  subreports,
		Status:      status,
	}
}
