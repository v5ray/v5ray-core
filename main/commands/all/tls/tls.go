package tls

import (
	"github.com/v4fly/v4ray-core/v0/main/commands/base"
)

// CmdTLS holds all tls sub commands
var CmdTLS = &base.Command{
	UsageLine: "{{.Exec}} tls",
	Short:     "TLS tools",
	Long: `{{.Exec}} {{.LongName}} provides tools for TLS.
	`,

	Commands: []*base.Command{
		cmdCert,
		cmdPing,
		cmdChainHash,
	},
}
