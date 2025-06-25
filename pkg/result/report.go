package result

type Status string

const (
	StatusSuccess Status = "success"
	// The verifier discovered an issue.
	StatusFailure Status = "failure"
	StatusSkipped Status = "skipped"
	// Unable to run the verifier, e.g., due to missing dependencies or configuration issues.
	StatusError Status = "error"
)

type Report struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      Status   `json:"status"`
	Output      string   `json:"output"`
	SubReports  []Report `json:"sub_reports,omitempty"`
}
