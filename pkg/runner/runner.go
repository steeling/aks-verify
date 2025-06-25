package runner

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/steeling/aks-verify/pkg/printer"
	"github.com/steeling/aks-verify/pkg/result"
	"github.com/steeling/aks-verify/pkg/verifier"
	"golang.org/x/sync/errgroup"
)

type Runner struct {
	verifiers []verifier.Verifier
}

func New(v []verifier.Verifier) *Runner {
	return &Runner{
		verifiers: v,
	}
}

func (r *Runner) Run(ctx context.Context, w io.Writer) error {
	g, ctx := errgroup.WithContext(ctx)
	var mu sync.Mutex
	reports := make([]result.Report, 0, len(r.verifiers))

	for _, v := range r.verifiers {
		v := v // capture range variable
		g.Go(func() error {
			_, err := fmt.Fprintln(w, "Running verifier:", v.Name())
			if err != nil {
				return err
			}
			report := v.Run(ctx)
			_, err = fmt.Fprintln(w, "Completed verifier:", v.Name(), "Status:", report.Status)
			if err != nil {
				return err
			}
			report.Name = v.Name()
			report.Description = v.Description()
			mu.Lock()
			reports = append(reports, report)

			mu.Unlock()
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		return fmt.Errorf("error running verifiers: %w", err)
	}
	p := printer.NewDefaultPrinter()
	return p.Print(w, reports)
}
