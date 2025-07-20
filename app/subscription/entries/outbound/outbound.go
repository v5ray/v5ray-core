package outbound

import (
	"github.com/v4fly/v4ray-core/v0/app/subscription/entries"
	"github.com/v4fly/v4ray-core/v0/app/subscription/specs"
	"github.com/v4fly/v4ray-core/v0/common"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

// NewOutboundEntriesParser internal api
func NewOutboundEntriesParser() entries.Converter {
	return newOutboundEntriesParser()
}

func newOutboundEntriesParser() entries.Converter {
	return &outboundEntriesParser{}
}

type outboundEntriesParser struct{}

func (o *outboundEntriesParser) ConvertToAbstractServerConfig(rawConfig []byte, kindHint string) (*specs.SubscriptionServerConfig, error) {
	parser := specs.NewOutboundParser()
	outbound, err := parser.ParseOutboundConfig(rawConfig)
	if err != nil {
		return nil, newError("failed to parse outbound config").Base(err).AtWarning()
	}
	return parser.ToSubscriptionServerConfig(outbound)
}

func init() {
	common.Must(entries.RegisterConverter("outbound", newOutboundEntriesParser()))
}
