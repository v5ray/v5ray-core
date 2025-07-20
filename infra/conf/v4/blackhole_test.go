package v4_test

import (
	"testing"

	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon"
	"github.com/v4fly/v4ray-core/v0/infra/conf/cfgcommon/testassist"
	v4 "github.com/v4fly/v4ray-core/v0/infra/conf/v4"
	"github.com/v4fly/v4ray-core/v0/proxy/blackhole"
)

func TestHTTPResponseJSON(t *testing.T) {
	creator := func() cfgcommon.Buildable {
		return new(v4.BlackholeConfig)
	}

	testassist.RunMultiTestCase(t, []testassist.TestCase{
		{
			Input: `{
				"response": {
					"type": "http"
				}
			}`,
			Parser: testassist.LoadJSON(creator),
			Output: &blackhole.Config{
				Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
			},
		},
		{
			Input:  `{}`,
			Parser: testassist.LoadJSON(creator),
			Output: &blackhole.Config{},
		},
	})
}
