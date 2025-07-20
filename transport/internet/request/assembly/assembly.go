package assembly

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/common"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

const protocolName = "request"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, newError("request is a transport protocol.")
	}))
}
