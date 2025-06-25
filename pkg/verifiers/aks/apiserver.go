package aks

// import controller runtime client
import (
	"context"

	"github.com/steeling/aks-verify/pkg/result"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type APIServerVerifier struct {
	client client.Client
}

func NewAPIServerVerifier(c client.Client) *APIServerVerifier {
	return &APIServerVerifier{
		client: c,
	}
}

func (v *APIServerVerifier) Name() string {
	return "API Server Configuration Verifier"
}

func (v *APIServerVerifier) Description() string {
	return "Verifies the API server configuration for best practices."
}

func (v *APIServerVerifier) Run(ctx context.Context) result.Report {
	// Implement the logic to verify the API server configuration
	// This is a placeholder implementation
	report := result.Report{
		Name:        "API Server Configuration",
		Description: "Verifies the API server configuration for best practices.",
		Status:      result.StatusSuccess,
	}

	// Example check: Verify if the API server is reachable
	if err := v.client.Get(ctx, client.ObjectKey{Name: "kube-system"}, &corev1.Namespace{}); err != nil {
		report.Status = result.StatusFailure
		report.Output = "API server is not reachable: " + err.Error()
	} else {
		report.Output = "API server is reachable."
	}

	return report
}
