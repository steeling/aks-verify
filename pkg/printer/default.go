package printer

import (
	"io"

	"github.com/steeling/aks-verify/pkg/result"
)

type Printer interface {
	Print([]result.Report) error
}

func NewDefaultPrinter() *DefaultPrinter {
	return &DefaultPrinter{}
}

type DefaultPrinter struct{}

func (p *DefaultPrinter) Print(w io.Writer, reports []result.Report) error {
	for _, report := range reports {
		if err := p.printReport(w, report, 0); err != nil {
			return err
		}
	}
	return nil
}

func (p *DefaultPrinter) printReport(w io.Writer, report result.Report, indent int) error {
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += "  "
	}

	status := string(report.Status)
	switch report.Status {
	case result.StatusSuccess:
		status = "\033[32m" + status + "\033[0m" // Green for success
	case result.StatusFailure:
		status = "\033[31m" + status + "\033[0m" // Red for failure
	case result.StatusSkipped:
		status = "\033[33m" + status + "\033[0m" // Yellow for skipped
	case result.StatusError:
		status = "\033[33m" + status + "\033[0m" // Yellow for error
	}

	// Print the report name, status, and description to writer
	_, err := io.WriteString(w, indentation+report.Name+" ("+status+")\n")
	if err != nil {
		return err
	}
	if report.Description != "" {
		_, err = io.WriteString(w, indentation+"  Description: "+report.Description+"\n")
		if err != nil {
			return err
		}
	}
	if report.Output != "" {
		_, err = io.WriteString(w, indentation+"  Output: "+report.Output+"\n")
		if err != nil {
			return err
		}
	}
	for _, subReport := range report.SubReports {
		if err := p.printReport(w, subReport, indent+1); err != nil {
			return err
		}
	}
	return nil
}
