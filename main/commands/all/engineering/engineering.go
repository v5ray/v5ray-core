package engineering

import "github.com/v4fly/v4ray-core/v0/main/commands/base"

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

var cmdEngineering = &base.Command{
	UsageLine: "{{.Exec}} engineering",
	Commands: []*base.Command{
		cmdConvertPb,
		cmdReversePb,
		cmdNonNativeLinkExtract,
		cmdNonNativeLinkExec,
		cmdSubscriptionEntriesExtract,
		cmdEncodeDataURL,
	},
}

func init() {
	base.RegisterCommand(cmdEngineering)
}
