package entries

import "github.com/v4fly/v4ray-core/v0/app/subscription/specs"

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type Converter interface {
	ConvertToAbstractServerConfig(rawConfig []byte, kindHint string) (*specs.SubscriptionServerConfig, error)
}
