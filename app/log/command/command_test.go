package command_test

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/v4fly/v4ray-core/v0"
	"github.com/v4fly/v4ray-core/v0/app/dispatcher"
	"github.com/v4fly/v4ray-core/v0/app/log"
	. "github.com/v4fly/v4ray-core/v0/app/log/command"
	"github.com/v4fly/v4ray-core/v0/app/proxyman"
	_ "github.com/v4fly/v4ray-core/v0/app/proxyman/inbound"
	_ "github.com/v4fly/v4ray-core/v0/app/proxyman/outbound"
	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/common/serial"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
