package v4

import (
	"github.com/golang/protobuf/proto"

	"github.com/v4fly/v4ray-core/v0/proxy/loopback"
)

type LoopbackConfig struct {
	InboundTag string `json:"inboundTag"`
}

func (l LoopbackConfig) Build() (proto.Message, error) {
	return &loopback.Config{InboundTag: l.InboundTag}, nil
}
