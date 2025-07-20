package dtls

import (
	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/transport/internet"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

const protocolName = "dtls"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
