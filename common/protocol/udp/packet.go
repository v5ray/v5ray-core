package udp

import (
	"github.com/v4fly/v4ray-core/v0/common/buf"
	"github.com/v4fly/v4ray-core/v0/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
