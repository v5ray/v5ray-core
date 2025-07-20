package tagged

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
