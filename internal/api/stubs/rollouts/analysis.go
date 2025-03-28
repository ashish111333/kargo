package rollouts

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	rolloutsapi "github.com/akuity/kargo/api/stubs/rollouts/v1alpha1"
)

// GetAnalysisTemplate returns a pointer to the AnalysisTemplate resource
// specified by the namespacedName argument. If no such resource is found, nil
// is returned instead.
func GetAnalysisTemplate(
	ctx context.Context,
	c client.Client,
	namespacedName types.NamespacedName,
) (*rolloutsapi.AnalysisTemplate, error) {
	at := rolloutsapi.AnalysisTemplate{}
	if err := c.Get(ctx, namespacedName, &at); err != nil {
		if err = client.IgnoreNotFound(err); err == nil {
			return nil, nil
		}
		return nil, fmt.Errorf(
			"error getting AnalysisTemplate %q in namespace %q: %w",
			namespacedName.Name,
			namespacedName.Namespace,
			err,
		)
	}
	return &at, nil
}

// GetClusterAnalysisTemplate returns a pointer to the ClusterAnalysisTemplate resource
// specified by the name argument. If no such resource is found, nil
// is returned instead.
func GetClusterAnalysisTemplate(
	ctx context.Context,
	c client.Client,
	name string,
) (*rolloutsapi.ClusterAnalysisTemplate, error) {
	cat := rolloutsapi.ClusterAnalysisTemplate{}
	if err := c.Get(
		ctx,
		types.NamespacedName{
			Name: name,
		},
		&cat,
	); err != nil {
		if err = client.IgnoreNotFound(err); err == nil {
			return nil, nil
		}
		return nil, fmt.Errorf(
			"error getting ClusterAnalysisTemplate %q: %w",
			name,
			err,
		)
	}
	return &cat, nil
}

func GetAnalysisRun(
	ctx context.Context,
	c client.Client,
	namespacedName types.NamespacedName,
) (*rolloutsapi.AnalysisRun, error) {
	ar := rolloutsapi.AnalysisRun{}
	if err := c.Get(ctx, namespacedName, &ar); err != nil {
		if err = client.IgnoreNotFound(err); err == nil {
			return nil, nil
		}
		return nil, fmt.Errorf(
			"error getting AnalysisRun %q in namespace %q: %w",
			namespacedName.Name,
			namespacedName.Namespace,
			err,
		)
	}
	return &ar, nil
}
