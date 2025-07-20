package udp

import (
	"context"
	"io"

	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/common/buf"
	"github.com/v4fly/v4ray-core/v0/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}

var DispatcherConnectionTerminationSignalReceiverMark = "DispatcherConnectionTerminationSignalReceiverMark"

type DispatcherConnectionTerminationSignalReceiver interface {
	io.Closer
}
