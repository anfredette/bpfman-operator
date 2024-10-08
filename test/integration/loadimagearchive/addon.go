//go:build integration_tests
// +build integration_tests

package loadimagearchive

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/kong/kubernetes-testing-framework/pkg/clusters"
	"github.com/kong/kubernetes-testing-framework/pkg/clusters/types/kind"
)

// -----------------------------------------------------------------------------
// LoadImageArchive Addon
// -----------------------------------------------------------------------------

const (
	// AddonName indicates the unique name of this addon.
	AddonName clusters.AddonName = "loadimagearchive"
)

type Addon struct {
	images []string
	ociBin string
	loaded bool
}

func New(ociBin string) clusters.Addon {
	return &Addon{ociBin: ociBin}
}

// -----------------------------------------------------------------------------
// LoadImageArchive Addon - Addon Implementation
// -----------------------------------------------------------------------------

func (a *Addon) Name() clusters.AddonName {
	return AddonName
}

func (a *Addon) Dependencies(_ context.Context, _ clusters.Cluster) []clusters.AddonName {
	return nil
}

func (a *Addon) Deploy(ctx context.Context, cluster clusters.Cluster) error {
	switch ctype := cluster.Type(); ctype {
	case kind.KindClusterType:
		if err := a.loadIntoKind(ctx, cluster); err != nil {
			return err
		}
	default:
		return fmt.Errorf("loadimagearchive addon is not supported by cluster type '%v'", cluster.Type())
	}

	return nil
}

func (a *Addon) Delete(_ context.Context, cluster clusters.Cluster) error {
	switch ctype := cluster.Type(); ctype {
	case kind.KindClusterType:
		// per https://github.com/kubernetes-sigs/kind/issues/658 this is basically impossible
		// we lie here, because we want to mask this error. not deleting an image from KIND is benign:
		// you either don't use it after (in which case you shouldn't care that it's still present) or
		// load another image with the same name (in which case the name will point to the new image)
		return nil
	default:
		return fmt.Errorf("loadimagearchive addon is not supported by cluster type '%v'", cluster.Type())
	}
}

func (a *Addon) Ready(context.Context, clusters.Cluster) ([]runtime.Object, bool, error) {
	// no way to verify this, we rely on Deploy's cmd.Run() not failing
	return nil, a.loaded, nil
}

func (a *Addon) DumpDiagnostics(context.Context, clusters.Cluster) (map[string][]byte, error) {
	diagnostics := make(map[string][]byte)
	return diagnostics, nil
}
