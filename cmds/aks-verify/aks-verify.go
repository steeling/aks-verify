package main

import (
	"context"
	"flag"
	"os"

	"github.com/steeling/aks-verify/pkg/runner"
	"github.com/steeling/aks-verify/pkg/verifier"
	"github.com/steeling/aks-verify/pkg/verifiers/aks"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	verifiers, err := setupVerifiers()
	if err != nil {
		panic(err)
	}

	runner := runner.New(verifiers)
	w := os.Stdout
	if err := runner.Run(ctx, w); err != nil {
		panic(err)
	}
}

func setupVerifiers() ([]verifier.Verifier, error) {
	k8sClient, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		return nil, err
	}

	apiserver := aks.NewAPIServerVerifier(k8sClient)

	return []verifier.Verifier{apiserver}, nil
}
