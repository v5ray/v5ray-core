package v4_test

import (
	"testing"

	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/common/protocol"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon/testassist"
	v4 "github.com/v4fly/v4ray-core/v0/infra/conf/v4"
	"github.com/v4fly/v4ray-core/v0/proxy/freedom"
)

func TestFreedomConfig(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.FreedomConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"domainStrategy": "AsIs",
				"timeout": 10,
				"redirect": "127.0.0.1:3366",
				"userLevel": 1
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &freedom.Config{
				DomainStrategy: freedom.Config_AS_IS,
				Timeout:        10,
				DestinationOverride: &freedom.DestinationOverride{
					Server: &protocol.ServerEndpoint{
						Address: &net.IPOrDomain{
							Address: &net.IPOrDomain_Ip{
								Ip: []byte{127, 0, 0, 1},
							},
						},
						Port: 3366,
					},
				},
				UserLevel: 1,
			},
		},
	})
}
