package v4_test

import (
	"testing"

	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon/testassist"
	v4 "github.com/v4fly/v4ray-core/v0/infra/conf/v4"
	"github.com/v4fly/v4ray-core/v0/proxy/dns"
)

func TestDnsProxyConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.DNSOutboundConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp"
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &dns.Config{
				Server: &net.Endpoint{
					Network: net.Network_TCP,
					Address: net.NewIPOrDomain(net.IPAddress([]byte{8, 8, 8, 8})),
					Port:    53,
				},
			},
		},
	})
}
