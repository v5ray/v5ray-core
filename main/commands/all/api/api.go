package api

import (
	"github.com/v4fly/v4ray-core/v0/main/commands/base"
)

// CmdAPI calls an API in an V2Ray process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "call V2Ray API",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate V2Ray via its API.
`,
	Commands: []*base.Command{
		cmdLog,
		cmdStats,
		cmdBalancerInfo,
		cmdBalancerOverride,
	},
}
