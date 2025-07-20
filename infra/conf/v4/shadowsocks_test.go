package v4_test

import (
	"testing"

	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/common/protocol"
	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon/testassist"
	v4 "github.com/v4fly/v4ray-core/v0/infra/conf/v4"
	"github.com/v4fly/v4ray-core/v0/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.ShadowsocksServerConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "v2ray-password"
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				User: &protocol.User{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "v2ray-password",
					}),
				},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
