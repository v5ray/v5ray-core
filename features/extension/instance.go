package extension

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/features"
)

// InstanceManagement : unstable
type InstanceManagement interface {
	features.Feature
	ListInstance(ctx context.Context) ([]string, error)
	AddInstance(ctx context.Context, name string, config []byte, configType string) error
	StartInstance(ctx context.Context, name string) error
	StopInstance(ctx context.Context, name string) error
	UntrackInstance(ctx context.Context, name string) error
}

func InstanceManagementType() interface{} {
	return (*InstanceManagement)(nil)
}
