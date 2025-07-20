package udp

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/common/buf"
	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/common/net/packetaddr"
	"github.com/v4fly/v4ray-core/v0/common/protocol/udp"
	"github.com/v4fly/v4ray-core/v0/features/routing"
)

type PacketAddrDispatcher struct {
	conn     net.PacketConn
	callback ResponseCallback
	ctx      context.Context
}

func (p PacketAddrDispatcher) Close() error {
	if p.ctx.Value(DispatcherConnectionTerminationSignalReceiverMark) != nil {
		p.ctx.Value(DispatcherConnectionTerminationSignalReceiverMark).(DispatcherConnectionTerminationSignalReceiver).Close()
	}
	return p.conn.Close()
}

func (p PacketAddrDispatcher) Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer) {
	if destination.Network != net.Network_UDP {
		return
	}

	// Processing of domain address is unsupported as it adds unpredictable overhead, it will be dropped.
	if destination.Address.Family().IsDomain() {
		return
	}

	p.conn.WriteTo(payload.Bytes(), &net.UDPAddr{IP: destination.Address.IP(), Port: int(destination.Port.Value())})
}

func (p PacketAddrDispatcher) readWorker() {
	for {
		readBuf := buf.New()
		n, addr, err := p.conn.ReadFrom(readBuf.Extend(2048))
		if err != nil {
			return
		}
		readBuf.Resize(0, int32(n))
		p.callback(p.ctx, &udp.Packet{Payload: readBuf, Source: net.DestinationFromAddr(addr)})
	}
}

type PacketAddrDispatcherCreator struct {
	ctx context.Context
}

func NewPacketAddrDispatcherCreator(ctx context.Context) PacketAddrDispatcherCreator {
	return PacketAddrDispatcherCreator{ctx: ctx}
}

func (pdc *PacketAddrDispatcherCreator) NewPacketAddrDispatcher(
	dispatcher routing.Dispatcher, callback ResponseCallback,
) DispatcherI {
	packetConn, _ := packetaddr.CreatePacketAddrConn(pdc.ctx, dispatcher, false)
	pd := &PacketAddrDispatcher{conn: packetConn, callback: callback, ctx: pdc.ctx}
	go pd.readWorker()
	return pd
}
