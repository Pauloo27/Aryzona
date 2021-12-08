package sysmon

import (
	"runtime"

	"github.com/Pauloo27/aryzona/command"
	"github.com/Pauloo27/aryzona/discord"
	"github.com/Pauloo27/aryzona/utils"
)

var Sys = command.Command{
	Name:        "sys",
	Description: "Show system info",
	Handler: func(ctx *command.CommandContext) {
		ctx.SuccessEmbed(
			discord.NewEmbed().
				WithTitle("System info").
				WithDescription(
					utils.Fmt(":computer: %s %s %s",
						runtime.GOOS, runtime.GOARCH, runtime.Version(),
					)),
		)
	},
}
