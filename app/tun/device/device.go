package device

import (
	"gvisor.dev/gvisor/pkg/tcpip/stack"

	"github.com/v4fly/v4ray-core/v0/common"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type Device interface {
	stack.LinkEndpoint

	common.Closable
}

type Options struct {
	Name string
	MTU  uint32
}

type DeviceConstructor func(Options) (Device, error)
